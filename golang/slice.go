package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	newElements := []int{}
	for i := range s {
		newElements = append(newElements, i)
	}
	s = append(s, newElements...)
	fmt.Println(s)
}
