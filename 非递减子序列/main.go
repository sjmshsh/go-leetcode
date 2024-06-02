package main

import "fmt"

// 使用回溯算法求出所有的子序列

var (
	path []int
	res  [][]int
)

func findSubsequences(nums []int) [][]int {
	path = make([]int, 0, len(nums))
	res = make([][]int, 0)
	dfs(nums, 0)
	return res
}

func dfs(nums []int, startIdx int) {
	if len(path) >= 2 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}

	// 初始化一个used map, 用来对同一层进行去重
	used := make(map[int]bool, len(nums))
	for i := startIdx; i < len(nums); i++ {
		if used[nums[i]] { // 去重
			continue
		}
		if len(path) == 0 || nums[i] >= path[len(path)-1] {
			used[nums[i]] = true
			path = append(path, nums[i])
			dfs(nums, i+1)
			path = path[:len(path)-1]
			// used[nums[i]] = false  这里注意, 这种情况是同层去重, 不能这么操作
		}
	}
}

func main() {
	res := findSubsequences([]int{4, 6, 7, 7})
	fmt.Println(res)
}
