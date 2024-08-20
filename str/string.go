package main

import "fmt"

func main() {
	s := "hello world"
	fmt.Println(s)
	fmt.Println(reverse(s))
}

func reverse(s string) string {
	s1 := []rune(s)
	for i := 0; i < len(s1)/2; i++ {
		s1[i], s1[len(s1)-1-i] = s1[len(s1)-1-i], s1[i]
	}
	return string(s1)
}
