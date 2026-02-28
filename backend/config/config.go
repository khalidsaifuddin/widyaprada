package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBType string `envconfig:"DB_TYPE" default:"sqlite"`

	HTTPPort    string `envconfig:"HTTP_PORT" default:"8080"`
	Environment string `envconfig:"ENVIRONMENT" default:"development"`

	// PostgreSQL
	Host     string `envconfig:"PGSQL_HOST" default:""`
	Port     string `envconfig:"PGSQL_PORT" default:"5432"`
	Username string `envconfig:"PGSQL_USERNAME" default:""`
	Password string `envconfig:"PGSQL_PASSWORD" default:""`
	DBName   string `envconfig:"PGSQL_DBNAME" default:""`

	LogMode         bool `envconfig:"DB_LOG_MODE" default:"true"`
	MaxIdleConns    int  `envconfig:"DB_MAX_IDLE_CONNS" default:"5"`
	MaxOpenConns    int  `envconfig:"DB_MAX_OPEN_CONNS" default:"10"`
	ConnMaxLifetime int  `envconfig:"DB_CONN_MAX_LIFETIME" default:"10"`

	// Redis (optional)
	RedisHost     string `envconfig:"REDIS_HOST" default:"127.0.0.1"`
	RedisPort     string `envconfig:"REDIS_PORT" default:"6379"`
	RedisPassword string `envconfig:"REDIS_PASSWORD" default:""`
	RedisMaxIdle  int    `envconfig:"REDIS_MAX_IDLE" default:"10"`
	RedisDBIndex  int    `envconfig:"REDIS_DB" default:"0"`

	// SQLite
	SQLiteDBPath string `envconfig:"SQLITE_DB_PATH" default:"./data/backend.db"`

	// Sentry (optional)
	IsEnableSentry         bool   `envconfig:"IS_ENABLE_SENTRY" default:"false"`
	SentryDSN              string `envconfig:"SENTRY_DSN" default:""`
	SentryTracesSampleRate string `envconfig:"SENTRY_TRACES_SAMPLE_RATE" default:"1.0"`
	SentryEnvironment      string `envconfig:"SENTRY_ENVIRONMENT" default:"development"`
	SentryRelease          string `envconfig:"SENTRY_RELEASE" default:"1.0.0"`

	// Swagger
	SwaggerHost     string `envconfig:"SWAGGER_HOST" default:"localhost:8080"`
	SwaggerBasePath string `envconfig:"SWAGGER_BASE_PATH" default:""`

	// JWT (Auth)
	JWTSecret   string `envconfig:"JWT_SECRET" default:"widyaprada-secret-change-in-production"`
	JWTExpiryHr int    `envconfig:"JWT_EXPIRY_HOUR" default:"1"`

	// Lupa Password
	FrontendURL        string `envconfig:"FRONTEND_URL" default:"http://localhost:3000"`
	ResetTokenExpiryHr int    `envconfig:"RESET_TOKEN_EXPIRY_HOUR" default:"1"`
}

func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}

func (c *Config) GetPostgreSQLConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.Username, c.Password, c.DBName)
}

func (c *Config) GetSQLiteDBPath() string {
	return c.SQLiteDBPath
}
