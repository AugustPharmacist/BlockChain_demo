package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	str1 := "helloworld"

	for i := 0; i < 1000000; i++ {
		hash := sha256.Sum256([]byte(str1 + string(i)))
		fmt.Printf("hash: %x, %d\n", hash[:], i)
	}
}
