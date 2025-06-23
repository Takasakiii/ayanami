package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func Encrypt(pass string, data []byte) ([]byte, error) {
	key := deriveKey(pass)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("crypt encrypt: cipher: %v", err)
	}

	gmc, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("crypt encrypt: gcm: %v", err)
	}

	nonce := make([]byte, gmc.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("crypt encrypt: nonce: %v", err)
	}

	finalData := gmc.Seal(nonce, nonce, data, nil)
	return finalData, nil
}
