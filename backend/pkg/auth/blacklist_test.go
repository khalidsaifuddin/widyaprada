package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMemoryBlacklist(t *testing.T) {
	bl := NewMemoryBlacklist()

	t.Run("Add and Contains", func(t *testing.T) {
		exp := time.Now().UTC().Add(time.Hour)
		bl.Add("token-123", exp)
		assert.True(t, bl.Contains("token-123"))
		assert.False(t, bl.Contains("unknown"))
	})

	t.Run("expired token not contained", func(t *testing.T) {
		exp := time.Now().UTC().Add(-time.Minute) // already expired
		bl.Add("expired-token", exp)
		assert.False(t, bl.Contains("expired-token"))
	})

	t.Run("new instance is empty", func(t *testing.T) {
		bl2 := NewMemoryBlacklist()
		assert.False(t, bl2.Contains("token-123"))
	})
}
