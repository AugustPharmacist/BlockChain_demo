package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	str1 := []string{"hello", "world", "!!!"}

	result := strings.Join(str1, "--")
	fmt.Printf("result: %s\n", result)

	//bytes.Join
	bytesArray := bytes.Join([][]byte{[]byte("hello"), []byte("world"), []byte("!!!")}, []byte("=="))
	fmt.Printf("bytesArray: %s\n", string(bytesArray))
}
