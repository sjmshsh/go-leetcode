package main

import (
	"fmt"
)

var (
	count1 int // count1计数器
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
				count1 = 0 // 重置计数器
				bfs(grid, visited, dirs, i, j)
				result = max(result, count1)
			}
		}
	}

	fmt.Println(result)
}

func bfs(grid [][]int, visited [][]bool, dirs [][]int, x, y int) {
	// 定义一个队列
	queue := make([][]int, 0)
	queue = append(queue, []int{x, y})
	visited[x][y] = true
	count1++

	for len(queue) > 0 {
		// 从队列里面拿到当前的值
		cur := queue[0]
		queue = queue[1:]

		curx := cur[0]
		cury := cur[1]
		// 遍历4个方向
		for _, d := range dirs {
			nx := curx + d[0]
			ny := cury + d[1]
			// 判断下一个位置的边界条件
			if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[0]) {
				continue
			}
			if !visited[nx][ny] && grid[nx][ny] == 1 {
				count1++
				queue = append(queue, []int{nx, ny})
				visited[nx][ny] = true
			}
		}
	}
}
