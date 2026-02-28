package conn

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/ProjectWidyaprada/backend/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitSQLiteDB(cfg *config.Config) *gorm.DB {
	dbPath := cfg.GetSQLiteDBPath()
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		log.Fatalf("Failed to create database directory: %s", err)
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		log.Fatalf("SQLite: %s", err.Error())
	}
	rdb, _ := db.DB()
	rdb.SetMaxIdleConns(cfg.MaxIdleConns)
	rdb.SetMaxOpenConns(cfg.MaxOpenConns)
	rdb.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Minute)
	log.Printf("Connected to SQLite: %s", dbPath)
	return db
}

func DbCloseSQLite(db *gorm.DB) {
	rdb, err := db.DB()
	if err != nil {
		return
	}
	_ = rdb.Close()
}
