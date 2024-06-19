package main

import (
	"fmt"
)

// 遇到一个没有遍历过的节点陆地计数器就加一
// 然后把该节点陆地能遍历到的陆地都标记上
// 在遇到标记过的陆地节点和海洋节点的时候直接跳过
// 这样计数器就是最终岛屿的数量

func dfs1(grid [][]int, visited [][]bool, x, y int, dir [][]int) {
	if visited[x][y] || grid[x][y] == 0 {
		return
	}
	visited[x][y] = true
	for _, d := range dir {
		nx := x + d[0]
		ny := y + d[1]
		if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[0]) {
			continue
		}
		dfs1(grid, visited, nx, ny, dir)
	}
}

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

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	// 定义dfs的方向
	dir := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	result := 0

	for i := range grid {
		for j := range grid[i] {
			if !visited[i][j] && grid[i][j] == 1 {
				result++
				dfs1(grid, visited, i, j, dir)
			}
		}
	}

	fmt.Println(result)
}
