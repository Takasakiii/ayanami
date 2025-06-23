package crypt

import (
	"crypto/sha256"
)

func deriveKey(pass string) []byte {
	h := sha256.Sum256([]byte(pass))
	return h[:]
}
