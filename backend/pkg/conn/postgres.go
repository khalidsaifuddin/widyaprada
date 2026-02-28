package conn

import (
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/ProjectWidyaprada/backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitPostgresDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		log.Fatalf("postgres: %s", err.Error())
	}
	rdb, _ := db.DB()
	rdb.SetMaxIdleConns(cfg.MaxIdleConns)
	rdb.SetMaxOpenConns(cfg.MaxOpenConns)
	rdb.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Minute)
	log.Print("Successfully connected to PostgreSQL")
	return db
}

func DbClosePostgres(db *gorm.DB) {
	rdb, err := db.DB()
	if err != nil {
		return
	}
	_ = rdb.Close()
}
