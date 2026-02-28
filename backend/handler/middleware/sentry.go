package middleware

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func InitSentry(dsn, environment, tracesSampleRate, release string) error {
	if dsn == "" {
		log.Println("Sentry DSN not set, skipping")
		return nil
	}
	rate, err := strconv.ParseFloat(tracesSampleRate, 64)
	if err != nil {
		rate = 1.0
	}
	opts := sentry.ClientOptions{
		Dsn:              dsn,
		Environment:      environment,
		TracesSampleRate: rate,
		EnableTracing:   true,
		Debug:            environment == "development",
	}
	if release != "" {
		opts.Release = release
	}
	if err := sentry.Init(opts); err != nil {
		return err
	}
	log.Printf("Sentry initialized: %s", environment)
	return nil
}

func SentryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tx := sentry.StartTransaction(c.Request.Context(), c.Request.Method+" "+c.FullPath())
		defer tx.Finish()
		c.Request = c.Request.WithContext(tx.Context())
		c.Next()
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				sentry.CaptureException(e.Err)
			}
		}
	}
}

func SentryRecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				hub := sentry.GetHubFromContext(c.Request.Context())
				if hub == nil {
					hub = sentry.CurrentHub()
				}
				hub.RecoverWithContext(context.WithValue(c.Request.Context(), sentry.RequestContextKey, c.Request), err)
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}

func FlushSentry() {
	if !sentry.Flush(2 * time.Second) {
		log.Println("Sentry flush timeout")
	}
}
