package middleware

import (
	"net/http"
	"strings"

	"github.com/ProjectWidyaprada/backend/config"
	example_api "github.com/ProjectWidyaprada/backend/handler/api/example"
	example_usecase "github.com/ProjectWidyaprada/backend/core/usecase/example"
	examplerepo "github.com/ProjectWidyaprada/backend/repository/example-repo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(cfg config.Config, db *gorm.DB) (*gin.Engine, interface{}) {
	if strings.EqualFold(cfg.Environment, "production") {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(CORSMiddleware())
	router.Use(InputValidationMiddleware())
	if cfg.IsEnableSentry {
		router.Use(SentryRecoveryMiddleware())
		router.Use(SentryMiddleware())
	}
	router.MaxMultipartMemory = 10 << 20
	router.Use(func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 1<<20)
		c.Next()
	})

	// Auto-migrate example table (skeleton; replace with migrations in production)
	_ = db.AutoMigrate(&examplerepo.Example{})

	// Example feature (hexagon: repo -> usecase -> handler)
	exampleRepo := examplerepo.NewExampleRepo(db)
	exampleUsecase := example_usecase.NewExampleUsecase(exampleRepo)
	exampleHandler := example_api.NewExampleHTTPHandler(exampleUsecase)

	apiGroup := router.Group("/api")
	v1Group := apiGroup.Group("/v1")

	examplesGroup := v1Group.Group("/examples")
	examplesGroup.GET("", exampleHandler.GetExampleList)
	examplesGroup.GET("/:id", exampleHandler.GetExampleDetail)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "404", "message": "Page not found"})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Widyaprada Backend API",
			"version": "1.0",
			"endpoints": gin.H{
				"health":  "/_health",
				"swagger": "/swagger/index.html",
				"api":     "/api/v1",
			},
		})
	})

	router.GET("/_health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	router.GET("/_db-test", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Database OK", "db_type": cfg.DBType})
	})

	return router, nil
}
