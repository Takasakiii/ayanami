package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func Decrypt(pass string, data []byte) ([]byte, error) {
	key := deriveKey(pass)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("crypt decrypt: cipher:  %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("crypt decrypt: gcm: %v", err)
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, fmt.Errorf("crypt decrypt: cipher: ciphertext too short")
	}

	nonce, ct := data[:nonceSize], data[nonceSize:]
	resultData, err := gcm.Open(nil, nonce, ct, nil)
	if err != nil {
		return nil, fmt.Errorf("crypt decrypt: gcm: %v", err)
	}

	return resultData, nil
}
