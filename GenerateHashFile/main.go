package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	fmt.Print("Hello")
}

func GenerateHash(data []byte) []byte {
	hash := sha1.Sum(data)
	return hash[:]
}
