package main

import (
	"fmt"
)

var dirs = [][]int{
	{0, 1},  // 右
	{1, 0},  // 下
	{-1, 0}, // 上
	{0, -1}, // 左
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &grid[i][j])
		}
	}
	// 上下遍历岛屿
	for i := 0; i < n; i++ {
		if grid[i][0] == 1 {
			dfs(grid, i, 0)
		}
		if grid[i][m-1] == 1 {
			dfs(grid, i, m-1)
		}
	}
	// 左右遍历岛屿
	for i := 0; i < m; i++ {
		if grid[0][i] == 1 {
			dfs(grid, 0, i)
		}
		if grid[n-1][i] == 1 {
			dfs(grid, n-1, i)
		}
	}
	// 最后一次遍历岛屿
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
			}
			if grid[i][j] == 2 {
				grid[i][j] = 1
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m-1; j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Printf("%d\n", grid[i][m-1])
	}
}

func dfs(grid [][]int, x, y int) {
	grid[x][y] = 2
	for i := 0; i < len(dirs); i++ {
		nx := x + dirs[i][0]
		ny := y + dirs[i][1]
		if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[0]) {
			continue
		}
		if grid[nx][ny] == 0 || grid[nx][ny] == 2 {
			continue
		}
		dfs(grid, nx, ny)
	}
}
