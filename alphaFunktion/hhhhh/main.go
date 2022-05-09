package main

import (
	"crypto/rand"
	"fmt"
	"log"
)

func main() {

	data, err := generateRandomBytes(16)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}

func generateRandomBytes(n int) ([]byte, error) {

	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}
