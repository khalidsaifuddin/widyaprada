package middleware

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

func InputValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, param := range c.Params {
			if !isValidInput(param.Value) {
				helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid input", "Path parameter contains invalid characters")
				c.Abort()
				return
			}
		}
		for _, values := range c.Request.URL.Query() {
			for _, value := range values {
				if !isValidInput(value) {
					helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid input", "Query parameter contains invalid characters")
					c.Abort()
					return
				}
			}
		}
		c.Next()
	}
}

func isValidInput(input string) bool {
	sqlPatterns := []string{
		`(?i)(union|select|insert|update|delete|drop|create|alter|exec|execute)`,
		`(?i)(<script|javascript:|vbscript:|onload=|onerror=|onclick=)`,
	}
	for _, pattern := range sqlPatterns {
		if matched, _ := regexp.MatchString(pattern, input); matched {
			return false
		}
	}
	lower := strings.ToLower(input)
	for _, s := range []string{"<script", "javascript:", "onerror=", "onclick=", "<iframe"} {
		if strings.Contains(lower, s) {
			return false
		}
	}
	return true
}
