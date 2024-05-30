package main

import (
	"fmt"
	"sort"
)

var (
	res  [][]int
	path []int
)

func combinationSum(candidates []int, target int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0)
	// 排序, 为剪枝做准备
	sort.Ints(candidates)
	combinationSumCore(candidates, target, 0)
	return res
}

func combinationSumCore(candidates []int, target int, start int) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		// 如果已经知道下一层的sum会大于target, 就没有必要进入下一层
		path = append(path, candidates[i])
		target -= candidates[i]
		combinationSumCore(candidates, target, i)
		target += candidates[i]
		path = path[:len(path)-1]
	}
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	res := combinationSum(candidates, target)
	fmt.Println(res)
}
