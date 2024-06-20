package main

import (
	"fmt"
)

var (
	count int // count计数器
)

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := range grid[i] {
			fmt.Scanf("%d", &grid[i][j])
		}
	}

	// 初始化visited数组
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	// 定义四个方向, 使用dfs算法
	dirs := [][]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}

	var result int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				count = 0 // 重置计数器
				dfs(grid, visited, dirs, i, j)
				result = max(result, count)
			}
		}
	}

	fmt.Println(result)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dfs(grid [][]int, visited [][]bool, dirs [][]int, x, y int) {
	if visited[x][y] || grid[x][y] == 0 {
		return
	}

	// 当前节点已经被dfs遍历过了，做一个标记
	visited[x][y] = true
	count++

	// 遍历4个方向，进行深度优先的遍历
	for _, d := range dirs {
		nx := x + d[0]
		ny := y + d[1]
		// 判断节点的边界条件
		if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[0]) {
			continue
		}
		dfs(grid, visited, dirs, nx, ny)
	}
}
