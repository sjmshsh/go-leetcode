package main

import "fmt"

func main() {
	num := 1
	switch num {
	case 1:
		fmt.Println("This is case 1")
		fallthrough
	case 2:
		fmt.Println("This is case 2")
	}
}
