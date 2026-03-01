package auth

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{"valid password", "password123", false},
		{"empty password", "", false},
		{"long password within bcrypt limit", strings.Repeat("a", 72), false},
		{"unicode password", "パスワード123", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashPassword(tt.password)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.NotEmpty(t, got)
			assert.NotEqual(t, tt.password, got)
		})
	}
}

func TestVerifyPassword(t *testing.T) {
	hash, err := HashPassword("correctpassword")
	require.NoError(t, err)

	tests := []struct {
		name         string
		plain        string
		hashed       string
		wantValid    bool
	}{
		{"correct password", "correctpassword", hash, true},
		{"wrong password", "wrongpassword", hash, false},
		{"empty plain", "", hash, false},
		{"empty hash", "any", "", false},
		{"invalid hash", "correctpassword", "not-a-valid-bcrypt-hash", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VerifyPassword(tt.plain, tt.hashed)
			assert.Equal(t, tt.wantValid, got)
		})
	}
}
