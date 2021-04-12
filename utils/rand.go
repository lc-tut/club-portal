package utils

import (
	"crypto/rand"
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

func GenerateRand15() (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	lettersLen := len(letters)

	b := make([]byte, 15)
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	var result string
	for _, v := range b {
		result += string(letters[int(v)%lettersLen])
	}

	return result, nil
}
