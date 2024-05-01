package main

import "fmt"

func main() {
	var t int
	for {
		var sum int
		n, _ := fmt.Scan(&t)
		if n == 0 {
			break
		} else {
			a := make([]int, t)
			for i := 0; i < t; i++ {
				fmt.Scan(&a[i])
			}
			for i := 0; i < t; i++ {
				sum += a[i]
			}
		}
		fmt.Println(sum)
	}
}
