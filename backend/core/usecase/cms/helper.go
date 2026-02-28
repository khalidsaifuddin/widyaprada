package cms

import (
	"crypto/rand"
	"encoding/hex"
)

func randomSuffix() string {
	b := make([]byte, 4)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
