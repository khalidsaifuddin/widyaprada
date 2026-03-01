package auth

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestGetClaimsFromContext(t *testing.T) {
	t.Run("no claims", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/", nil)
		claims := GetClaimsFromContext(c)
		assert.Nil(t, claims)
	})

	t.Run("with claims", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/", nil)
		expected := &Claims{UserID: "user-1", Email: "a@b.com"}
		c.Set(ContextKeyClaims, expected)
		claims := GetClaimsFromContext(c)
		assert.Equal(t, expected, claims)
		assert.Equal(t, "user-1", claims.UserID)
	})

	t.Run("wrong type", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/", nil)
		c.Set(ContextKeyClaims, "not-claims")
		claims := GetClaimsFromContext(c)
		assert.Nil(t, claims)
	})
}

func TestGetTokenFromContext(t *testing.T) {
	t.Run("no token", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/", nil)
		token := GetTokenFromContext(c)
		assert.Empty(t, token)
	})

	t.Run("with token", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/", nil)
		c.Set(ContextKeyToken, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
		token := GetTokenFromContext(c)
		assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9", token)
	})
}
