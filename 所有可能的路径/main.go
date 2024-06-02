package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	result := make([][]int, 0)

	var dfs func(path []int, step int)
	dfs = func(path []int, step int) {
		// 从0遍历到length - 1
		if step == len(graph)-1 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		for i := 0; i < len(graph[step]); i++ {
			next := append(path, graph[step][i])
			dfs(next, graph[step][i])
		}
	}

	// 从0开始, 开始push 0进去
	dfs([]int{0}, 0)
	return result
}

func main() {
	res := allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}})
	fmt.Println(res)
}
