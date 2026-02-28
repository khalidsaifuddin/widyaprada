package middleware

import (
	"net/http"
	"strings"

	"github.com/ProjectWidyaprada/backend/config"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/gin-gonic/gin"
)

// AuthRequired middleware: validasi Bearer token, cek blacklist, set claims di context
func AuthRequired(cfg config.Config, blacklist *auth.MemoryBlacklist) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak ditemukan"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Format token tidak valid"})
			c.Abort()
			return
		}
		tokenString := strings.TrimSpace(parts[1])
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak ditemukan"})
			c.Abort()
			return
		}

		claims, err := auth.ParseToken(tokenString, cfg.JWTSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak valid atau kadaluarsa"})
			c.Abort()
			return
		}

		// Cek blacklist (logout)
		if blacklist.Contains(claims.ID) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token sudah logout"})
			c.Abort()
			return
		}

		c.Set(auth.ContextKeyClaims, claims)
		c.Set(auth.ContextKeyToken, tokenString)
		c.Next()
	}
}

// RequireSuperAdmin middleware: harus dipanggil setelah AuthRequired. Hanya Super Admin yang boleh akses.
func RequireSuperAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := auth.GetClaimsFromContext(c)
		if claims == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		for _, code := range claims.Roles {
			if code == "SUPER_ADMIN" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"message": "Hanya Super Admin yang dapat mengakses fitur ini"})
		c.Abort()
	}
}

// RequireAnyRole middleware: harus dipanggil setelah AuthRequired. User harus punya salah satu role.
func RequireAnyRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := auth.GetClaimsFromContext(c)
		if claims == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		for _, code := range claims.Roles {
			for _, allowed := range allowedRoles {
				if code == allowed {
					c.Next()
					return
				}
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"message": "Anda tidak memiliki akses untuk fitur ini"})
		c.Abort()
	}
}
