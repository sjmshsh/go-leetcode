package main

import (
	"fmt"
)

func main() {
	var b []byte

	for i := 0; i < 1000; i++ {
		b = append(b, "hello"...)
	}
	b = append(b, "Hello"...)
	b = append(b, ", "...)
	b = append(b, "World!"...)

	result := string(b)
	fmt.Println(result)
}
