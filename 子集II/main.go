package main

import (
	"fmt"
	"sort"
)

var (
	path []int
	res  [][]int
)

func subsetsWithDup(nums []int) [][]int {
	path = make([]int, 0, len(nums))
	res = make([][]int, 0)
	sort.Ints(nums)
	dfs(nums, 0)
	return res
}

func dfs(nums []int, startIdx int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	res = append(res, tmp)

	for i := startIdx; i < len(nums); i++ {
		if i != startIdx && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		dfs(nums, i+1)
		path = path[:len(path)-1]
	}
}

func main() {
	res := subsetsWithDup([]int{1, 2, 2})
	fmt.Println(res)
}
