package main

import "fmt"

// 思路：本题要求找到不靠边的陆地面积，那么我们只要从周边找到陆地然后
// 通过dfs或者bfs将周边靠陆地且相邻的陆地都变成海洋，然后再去重新遍历地图，
// 统计此时还剩下的陆地就可以了
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
	// 在上面和下面进行dfs遍历
	for i := 0; i < m; i++ {
		if grid[0][i] == 1 {
			bfs(grid, 0, i)
		}
		if grid[n-1][i] == 1 {
			bfs(grid, n-1, i)
		}
	}
	// 在左边和右边进行dfs遍历
	for i := 0; i < n; i++ {
		if grid[i][0] == 1 {
			bfs(grid, i, 0)
		}
		if grid[i][m-1] == 1 {
			bfs(grid, i, m-1)
		}
	}

	result = 0
	// 接下来再进行遍历即可
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				bfs(grid, i, j)
			}
		}
	}

	fmt.Println(result)
}

func bfs(grid [][]int, x, y int) {
	// 定义一个队列
	queue := make([][]int, 0)
	// 往队列里面插入当前的位置
	queue = append(queue, []int{x, y})
	// 只要加入队列就立即标记
	grid[x][y] = 0
	result++
	for len(queue) > 0 {
		// 取出队头元素
		cur := queue[0]
		queue = queue[1:]
		// 获取当前的坐标
		curx := cur[0]
		cury := cur[1]
		// 遍历4个方向, 只要加入队列
		for i := 0; i < len(dirs); i++ {
			nx := curx + dirs[i][0]
			ny := cury + dirs[i][1]
			if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[0]) {
				continue
			}
			if grid[nx][ny] == 0 {
				continue
			}
			grid[nx][ny] = 0
			queue = append(queue, []int{nx, ny})
		}
	}
}
