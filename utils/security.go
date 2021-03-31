package utils

import (
	"crypto/rand"
	"fmt"
)

func GenerateRand16() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
