package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func RandBytes(size int) (string, error) {
	b := make([]byte, size)
	_, err := rand.Read(b) // записываем байты в слайс b
	if err != nil {
		return "", fmt.Errorf("random read error: %w", err)
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func main() {
	fmt.Println(RandBytes(20))
}
