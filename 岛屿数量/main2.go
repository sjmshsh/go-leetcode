package main

import "fmt"

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

	// 定义visited数组
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	var result int

	// 定义广搜的方向
	dirs := [][]int{
		{0, 1},  // 右
		{1, 0},  // 下
		{-1, 0}, // 上
		{0, -1}, // 左
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 如果没有访问过该节点并且是陆地的话就进行bfs
			if !visited[i][j] && grid[i][j] == 1 {
				result++
				bfs(grid, visited, dirs, i, j)
			}
		}
	}

	fmt.Println(result)
}

func bfs(grid [][]int, visited [][]bool, dirs [][]int, x, y int) {
	// 定义一个队列
	queue := make([][]int, 0)
	// 将当前元素放入队列
	queue = append(queue, []int{x, y})
	// 标记
	visited[x][y] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		// 拿到当前的位置坐标
		curx := cur[0]
		cury := cur[1]
		// 遍历4个方向并加入到队列当中
		for i := 0; i < len(dirs); i++ {
			nx := curx + dirs[i][0]
			ny := cury + dirs[i][1]
			// 判断下一个位置的边界条件
			if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[0]) {
				continue
			}
			// 判断是否是陆地以及是否已经访问过了
			if grid[nx][ny] == 1 && !visited[nx][ny] {
				queue = append(queue, []int{nx, ny})
				visited[nx][ny] = true
			}
		}
	}
}
