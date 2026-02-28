package auth

import (
	"sync"
	"time"
)

// tokenEntry menyimpan token yang di-blacklist dengan expiry (untuk cleanup)
type tokenEntry struct {
	expiresAt time.Time
}

// MemoryBlacklist in-memory token blacklist (untuk single-instance)
// Untuk production multi-instance, gunakan Redis
type MemoryBlacklist struct {
	mu     sync.RWMutex
	tokens map[string]tokenEntry
}

// NewMemoryBlacklist membuat instance blacklist
func NewMemoryBlacklist() *MemoryBlacklist {
	return &MemoryBlacklist{tokens: make(map[string]tokenEntry)}
}

// Add menambahkan token ke blacklist (bisa raw token atau jti)
func (b *MemoryBlacklist) Add(key string, expiresAt time.Time) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.tokens[key] = tokenEntry{expiresAt: expiresAt}
}

// Contains mengecek apakah token ada di blacklist
// Token yang sudah expired dianggap tidak blacklist (bisa di-cleanup nanti)
func (b *MemoryBlacklist) Contains(key string) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	entry, ok := b.tokens[key]
	if !ok {
		return false
	}
	// Jika sudah expired, anggap tidak blacklist (token JWT sendiri sudah invalid)
	if time.Now().UTC().After(entry.expiresAt) {
		return false
	}
	return true
}
