package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&g[i][j])
			fmt.Print(g[i][j], " ")
		}
		fmt.Println()
	}
}
