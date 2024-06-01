package main

import "fmt"

var (
	path []int
	res  [][]int
)

func subsets(nums []int) [][]int {
	path = make([]int, 0, len(nums))
	res = make([][]int, 0)
	dfs(nums, 0)
	return res
}

// 子集求的是树的所有节点, 因此没有剪枝
func dfs(nums []int, start int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	res = append(res, tmp)

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		dfs(nums, i+1)
		path = path[:len(path)-1]
	}
}

func main() {
	res := subsets([]int{1, 2, 3})
	fmt.Println(res)
}
