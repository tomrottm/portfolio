package storage

import (
	"crypto/rand"
	"fmt"
)

// GenID returns a random ID string of format "prefix_random16chars",
func GenID(prefix string) (string, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s_%x", prefix, b), nil
}
