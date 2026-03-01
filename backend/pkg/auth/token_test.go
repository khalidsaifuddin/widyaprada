package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateResetToken(t *testing.T) {
	plain, hashed, err := GenerateResetToken()
	require.NoError(t, err)
	assert.Len(t, plain, 64) // 32 bytes = 64 hex chars
	assert.Len(t, hashed, 64)
	assert.NotEqual(t, plain, hashed)
}

func TestGenerateResetToken_Uniqueness(t *testing.T) {
	p1, h1, err := GenerateResetToken()
	require.NoError(t, err)
	p2, h2, err := GenerateResetToken()
	require.NoError(t, err)
	assert.NotEqual(t, p1, p2)
	assert.NotEqual(t, h1, h2)
}

func TestHashResetToken(t *testing.T) {
	plain, hashed, err := GenerateResetToken()
	require.NoError(t, err)

	computed := HashResetToken(plain)
	assert.Equal(t, hashed, computed)
}

func TestHashResetToken_Deterministic(t *testing.T) {
	plain := "abc123"
	assert.Equal(t, HashResetToken(plain), HashResetToken(plain))
}
