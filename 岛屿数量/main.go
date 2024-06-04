package main

import "fmt"

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, r, c int) {
	h, w := len(grid), len(grid[0])
	if r < 0 || r >= h || c < 0 || c >= w {
		return
	}
	if grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	res := numIslands(grid)
	fmt.Println(res)
}
