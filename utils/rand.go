package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GenerateCSRFState() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func GenerateSlug(uuid string) string {
	hashedUUID := sha256.Sum256([]byte(uuid))
	str := hex.EncodeToString(hashedUUID[:])
	result := str[:15]

	return result
}

func GenerateFileName(fn string) (string, error) {
	byteFn := []byte(fn)
	hashedFile := sha256.Sum256(byteFn)
	sliceHash := hashedFile[:5]

	b := make([]byte, 4)
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x%x", sliceHash, b), nil
}
