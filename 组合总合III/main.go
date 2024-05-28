package main

import "fmt"

var (
	path []int
	res  [][]int
)

func combinationSum3(k int, n int) [][]int {
	path = make([]int, 0, k)
	res = make([][]int, 0)
	dfs(k, n, 1, 0)
	return res
}

// k代表树的深度
// n代表target
func dfs(k, n, start, sum int) {
	if len(path) == k {
		if sum == n {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
	}
	for i := start; i <= 9; i++ {
		// 剪枝
		// 如果下一步的sum已经超过了target, 就没有必要进入到下一步
		if sum+i > n || 9-i+1 < k-len(path) {
			break
		}
		path = append(path, i)
		sum += i
		dfs(k, n, i+1, sum)
		path = path[:len(path)-1]
		sum -= i
	}
}

func main() {
	res := combinationSum3(3, 7)
	fmt.Println(res)
}
