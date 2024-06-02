package main

import "fmt"

var (
	path []int
	res  [][]int
	st   []bool // state的缩写
)

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0, len(nums))
	st = make([]bool, len(nums))
	dfs(nums, 0)
	return res
}

func dfs(nums []int, cur int) {
	if cur == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}

	for i := 0; i < len(nums); i++ {
		// 如果已经使用过，则跳过
		if !st[i] {
			st[i] = true
			path = append(path, nums[i])
			dfs(nums, cur+1)
			path = path[:len(path)-1]
			st[i] = false
		}
	}
}

func main() {
	res := permute([]int{1, 2, 3})
	fmt.Println(res)
}
