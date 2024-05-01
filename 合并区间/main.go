package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// 将区间按照左端点升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := make([][]int, 0)
	ans = append(ans, intervals[0])
	for _, e := range intervals[:1] {
		if ans[len(ans)-1][1] < e[0] {
			// 如果当前数组的左边大于ans最后一个数组的右边
			// 说明必定不会重合
			ans = append(ans, e)
		} else {
			// 说明会重合, 需要更新ans最后一个的值
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], e[1])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	res := merge(intervals)
	fmt.Println(res)
}
