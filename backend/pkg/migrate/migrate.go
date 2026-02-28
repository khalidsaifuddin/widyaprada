package migrate

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"

	"github.com/ProjectWidyaprada/backend/config"
	"gorm.io/gorm"
)

const (
	tableName = "schema_migrations"
)

// Migration represents a single migration file.
type Migration struct {
	Version  string // e.g. "001"
	Filename string // e.g. "001_create_ref_schema.sql"
}

// Run executes auto-migration: ensures schema_migrations table exists,
// detects current DB version, finds stale (pending) migrations, and runs them in order.
// Only runs when DB_TYPE=postgres; returns nil for sqlite (no-op).
func Run(db *gorm.DB, cfg *config.Config) error {
	if cfg.DBType != "postgres" {
		log.Print("migrate: skipped (PostgreSQL only)")
		return nil
	}

	m := &runner{db: db}
	if err := m.ensureTable(); err != nil {
		return fmt.Errorf("migrate: ensure schema_migrations table: %w", err)
	}

	applied, err := m.getAppliedVersions()
	if err != nil {
		return fmt.Errorf("migrate: get applied versions: %w", err)
	}

	available, err := m.scanMigrations(cfg.MigrationsPath)
	if err != nil {
		return fmt.Errorf("migrate: scan migrations: %w", err)
	}

	pending := m.findPending(applied, available)
	if len(pending) == 0 {
		log.Printf("migrate: database up to date (version %s)", m.latestVersion(applied))
		return nil
	}

	log.Printf("migrate: current version %s, running %d pending migration(s)", m.latestVersion(applied), len(pending))
	for _, mig := range pending {
		if err := m.runOne(mig, cfg.MigrationsPath); err != nil {
			return fmt.Errorf("migrate: run %s: %w", mig.Filename, err)
		}
		log.Printf("migrate: applied %s", mig.Filename)
	}
	log.Printf("migrate: done, database at version %s", pending[len(pending)-1].Version)
	return nil
}

type runner struct {
	db *gorm.DB
}

func (m *runner) ensureTable() error {
	sql := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS public.%s (
			version VARCHAR(32) PRIMARY KEY,
			filename VARCHAR(255) NOT NULL,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`, tableName)
	return m.db.Exec(sql).Error
}

func (m *runner) getAppliedVersions() (map[string]bool, error) {
	var rows []struct {
		Version string
	}
	err := m.db.Table(tableName).Select("version").Find(&rows).Error
	if err != nil {
		return nil, err
	}
	applied := make(map[string]bool)
	for _, r := range rows {
		applied[r.Version] = true
	}
	return applied, nil
}

// migrationFileRegex matches NNN_name.sql (e.g. 001_create_ref_schema.sql).
var migrationFileRegex = regexp.MustCompile(`^(\d{3})_(.+)\.sql$`)

// resolveMigrationsDir returns migrations dir. Uses configured path if it exists; otherwise tries "migrations", "../migrations".
func resolveMigrationsDir(configured string) string {
	candidates := []string{configured, "migrations", "../migrations"}
	for _, d := range candidates {
		if d == "" {
			continue
		}
		if fi, err := os.Stat(d); err == nil && fi.IsDir() {
			return d
		}
	}
	if configured != "" {
		return configured // use as-is; will fail with clear error when reading
	}
	return "migrations"
}

func (m *runner) scanMigrations(migrationsPath string) ([]Migration, error) {
	dir := resolveMigrationsDir(migrationsPath)
	return m.scanFromFS(dir)
}

func (m *runner) scanFromFS(dir string) ([]Migration, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read dir %q: %w", dir, err)
	}
	var migs []Migration
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		matches := migrationFileRegex.FindStringSubmatch(e.Name())
		if len(matches) < 3 {
			continue
		}
		migs = append(migs, Migration{Version: matches[1], Filename: e.Name()})
	}
	sort.Slice(migs, func(i, j int) bool { return migs[i].Version < migs[j].Version })
	return migs, nil
}

func (m *runner) findPending(applied map[string]bool, available []Migration) []Migration {
	var pending []Migration
	for _, mig := range available {
		if !applied[mig.Version] {
			pending = append(pending, mig)
		}
	}
	return pending
}

func (m *runner) latestVersion(applied map[string]bool) string {
	if len(applied) == 0 {
		return "000"
	}
	max := "000"
	for v := range applied {
		if v > max {
			max = v
		}
	}
	return max
}

func (m *runner) runOne(mig Migration, migrationsPath string) error {
	dir := resolveMigrationsDir(migrationsPath)
	path := filepath.Join(dir, mig.Filename)
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	tx := m.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Execute migration SQL (PostgreSQL supports multi-statement in one Exec)
	if err := tx.Exec(string(content)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("exec: %w", err)
	}

	// Record applied migration
	recordSQL := fmt.Sprintf(
		`INSERT INTO public.%s (version, filename, applied_at) VALUES ($1, $2, NOW())`,
		tableName,
	)
	if err := tx.Exec(recordSQL, mig.Version, mig.Filename).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("record migration: %w", err)
	}

	return tx.Commit().Error
}
