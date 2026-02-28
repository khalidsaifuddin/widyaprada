package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

// GenerateResetToken generates random 32-byte token (64 hex chars)
func GenerateResetToken() (plain, hashed string, err error) {
	b := make([]byte, 32)
	if _, err = rand.Read(b); err != nil {
		return "", "", err
	}
	plain = hex.EncodeToString(b)
	hash := sha256.Sum256([]byte(plain))
	hashed = hex.EncodeToString(hash[:])
	return plain, hashed, nil
}

// HashResetToken hashes a plain token for lookup
func HashResetToken(plain string) string {
	hash := sha256.Sum256([]byte(plain))
	return hex.EncodeToString(hash[:])
}
