package main

import (
	"fmt"
	"sort"
)

func main() {
	var w []int
	var x int
	for {
		_, err := fmt.Scan(&x)
		if err != nil {
			break
		}
		w = append(w, x)
	}
	sort.Ints(w)
	for _, x := range w {
		fmt.Print(x, " ")
	}
}
