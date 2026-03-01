package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateToken(t *testing.T) {
	secret := "test-secret-123"
	userID := "user-1"
	email := "user@example.com"
	roles := []string{"ADMIN", "USER"}
	expiryHours := 2

	token, expiresIn, err := GenerateToken(userID, email, roles, secret, expiryHours)
	require.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, expiryHours*3600, expiresIn)
}

func TestGenerateToken_EmptySecret(t *testing.T) {
	token, _, err := GenerateToken("u1", "e@x.com", nil, "", 1)
	require.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestGenerateToken_ZeroExpiry(t *testing.T) {
	token, expiresIn, err := GenerateToken("u1", "e@x.com", nil, "secret", 0)
	require.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, 3600, expiresIn) // defaults to 1 hour
}

func TestParseToken(t *testing.T) {
	secret := "test-secret-123"
	userID := "user-99"
	email := "test@example.com"
	roles := []string{"PESERTA"}

	token, _, err := GenerateToken(userID, email, roles, secret, 1)
	require.NoError(t, err)

	claims, err := ParseToken(token, secret)
	require.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, email, claims.Email)
	assert.Equal(t, roles, claims.Roles)
	assert.NotEmpty(t, claims.ID)
	assert.True(t, claims.ExpiresAt.Time.After(time.Now().UTC()))
}

func TestParseToken_InvalidSecret(t *testing.T) {
	secret := "secret"
	token, _, err := GenerateToken("u1", "e@x.com", nil, secret, 1)
	require.NoError(t, err)

	claims, err := ParseToken(token, "wrong-secret")
	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestParseToken_InvalidTokenString(t *testing.T) {
	claims, err := ParseToken("invalid.jwt.token", "secret")
	assert.Error(t, err)
	assert.Nil(t, claims)
}
