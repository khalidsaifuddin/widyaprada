package main

import (
	"fmt"
	"log"

	"github.com/ProjectWidyaprada/backend/config"
	"github.com/ProjectWidyaprada/backend/docs"
	"github.com/ProjectWidyaprada/backend/handler/middleware"
	"github.com/ProjectWidyaprada/backend/pkg/conn"
	"github.com/ProjectWidyaprada/backend/pkg/migrate"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// @title           Widyaprada Backend API
// @version         1.0
// @description     Backend service (hexagon/clean architecture).
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	log.Print("initializing backend")

	cfg := config.Get()

	if cfg.IsEnableSentry {
		if err := middleware.InitSentry(cfg.SentryDSN, cfg.SentryEnvironment, cfg.SentryTracesSampleRate, cfg.SentryRelease); err != nil {
			log.Printf("Failed to initialize Sentry: %v", err)
		}
	}
	if cfg.IsEnableSentry {
		defer middleware.FlushSentry()
	}

	var db *gorm.DB
	switch cfg.DBType {
	case "postgres":
		log.Print("Connecting to PostgreSQL...")
		db = conn.InitPostgresDB(&cfg)
		defer conn.DbClosePostgres(db)
	case "sqlite":
		log.Print("Connecting to SQLite...")
		db = conn.InitSQLiteDB(&cfg)
		defer conn.DbCloseSQLite(db)
	default:
		log.Fatalf("Unsupported DB_TYPE: %s (use postgres or sqlite)", cfg.DBType)
	}

	if err := migrate.Run(db, &cfg); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	router, _ := middleware.InitRouter(cfg, db)

	if cfg.SwaggerHost != "" {
		docs.SwaggerInfo.Host = cfg.SwaggerHost
	}
	if cfg.SwaggerBasePath != "" {
		docs.SwaggerInfo.BasePath = cfg.SwaggerBasePath
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(":" + cfg.HTTPPort); err != nil {
		panic(fmt.Errorf("failed to start server: %w", err))
	}
}
