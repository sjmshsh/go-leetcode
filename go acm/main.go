package main

import "fmt"

func main() {
	a := 0
	b := 0
	for {
		n, _ := fmt.Scanln(&a, &b)
		if n == 0 {
			break
		} else {
			fmt.Printf("%d\n", a+b)
		}
	}
}
