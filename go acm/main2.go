package main

import (
	"fmt"
)

func main() {
	var t int
	for {
		var sum int
		fmt.Scan(&t)
		if t == 0 {
			break
		} else {
			arr := make([]int, t)
			for i := 0; i < t; i++ {
				fmt.Scan(&arr[i])
			}
			for i := 0; i < t; i++ {
				sum += arr[i]
			}
		}
		fmt.Println(sum)
	}
}
