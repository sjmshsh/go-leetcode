package main

import (
	"fmt"
	"sort"
)

var (
	path []int
	res  [][]int
	used []bool
)

func combinationSum2(candidates []int, target int) [][]int {
	path = make([]int, 0, len(candidates))
	res = make([][]int, 0)
	used = make([]bool, len(candidates))
	sort.Ints(candidates)
	combinationSum2Core(candidates, target, 0)
	return res
}

func combinationSum2Core(candidates []int, target int, startIndex int) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := startIndex; i <= len(candidates); i++ {
		if target < candidates[i] {
			break
		}
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
			continue
		}
		used[i] = true
		path = append(path, candidates[i])
		target -= candidates[i]
		combinationSum2Core(candidates, target, i+1) // 因为一个数字只能使用一次
		target += candidates[i]
		path = path[:len(path)-1]
		used[i] = false
	}
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	res := combinationSum2(candidates, target)
	fmt.Println(res)
}
