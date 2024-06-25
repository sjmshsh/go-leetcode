package main

import "fmt"

var (
	result int
	dirs   = [][]int{
		{0, 1},  // 右
		{-1, 0}, // 上
		{1, 0},  // 下
		{0, -1}, // 左
	}
)

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
			dfs(grid, 0, i)
		}
		if grid[n-1][i] == 1 {
			dfs(grid, n-1, i)
		}
	}
	// 在左边和右边进行dfs遍历
	for i := 0; i < n; i++ {
		if grid[i][0] == 1 {
			dfs(grid, i, 0)
		}
		if grid[i][m-1] == 1 {
			dfs(grid, i, m-1)
		}
	}

	result = 0
	// 接下来再进行遍历即可
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				dfs(grid, i, j)
			}
		}
	}

	fmt.Println(result)
}

func dfs(grid [][]int, x, y int) {
	if grid[x][y] == 0 {
		return
	}

	// 变成水
	grid[x][y] = 0
	result++

	// 四个方向遍历
	for i := 0; i < len(dirs); i++ {
		nx := x + dirs[i][0]
		ny := y + dirs[i][1]
		// 不能超过边界
		if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[0]) {
			continue
		}
		// 不符合条件不继续进行遍历
		if grid[nx][ny] == 0 {
			continue
		}
		dfs(grid, nx, ny)
	}

}
