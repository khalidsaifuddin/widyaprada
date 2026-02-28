package auth

import (
	"github.com/gin-gonic/gin"
)

// Context keys untuk auth (dipakai middleware dan handler)
const (
	ContextKeyClaims = "claims"
	ContextKeyToken  = "token_raw"
)

// GetClaimsFromContext mengambil claims dari context (setelah AuthRequired)
func GetClaimsFromContext(c *gin.Context) *Claims {
	v, exists := c.Get(ContextKeyClaims)
	if !exists {
		return nil
	}
	claims, ok := v.(*Claims)
	if !ok {
		return nil
	}
	return claims
}

// GetTokenFromContext mengambil raw token dari context
func GetTokenFromContext(c *gin.Context) string {
	v, exists := c.Get(ContextKeyToken)
	if !exists {
		return ""
	}
	s, _ := v.(string)
	return s
}
