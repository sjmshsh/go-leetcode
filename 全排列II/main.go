package main

import (
	"fmt"
	"sort"
)

var (
	res  [][]int
	path []int
	st   []bool
)

func permuteUnique(nums []int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0, len(nums))
	st = make([]bool, len(nums))
	// 对数组进行排序, 方便去重复
	sort.Ints(nums)
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
		// 去重
		if i != 0 && nums[i] == nums[i-1] && !st[i-1] {
			continue
		}
		if !st[i] {
			path = append(path, nums[i])
			st[i] = true
			dfs(nums, cur+1)
			path = path[:len(path)-1]
			st[i] = false
		}
	}
}

func main() {
	res := permuteUnique([]int{1, 2, 3})
	fmt.Println(res)
}
