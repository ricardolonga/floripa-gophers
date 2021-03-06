package main

import (
	"bytes"
	"fmt"
)

func main() {
	// START OMIT
	text := []byte("Boa noite, galera!")
	fmt.Printf("text:  %s", desc(text))

	hello := bytes.Fields(text)[0]
	fmt.Printf("hello: %s", desc(hello))

	hello = append(hello, '#')
	fmt.Printf("hello: %s", desc(hello))

	hello[1] = 'Z'
	fmt.Printf("hello: %s", desc(hello))

	fmt.Printf("text:  %s", desc(text))
	// END OMIT
}

func desc(b []byte) string {
	return fmt.Sprintf("len: %2d | cap: %2d | %q\n", len(b), cap(b), b)
}