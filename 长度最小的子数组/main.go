package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	ans := math.MaxInt32
	// 窗口内数组的和
	sum := 0
	n := len(nums)
	i := 0 // 滑动窗口的起始位置
	// j可以理解为滑动窗口的右指针
	for j := 0; j < n; j++ {
		// 尝试扩大装口
		sum += nums[j]
		// 检查是否需要收缩窗口
		for sum >= target {
			len := j - i + 1 // 窗口的大小
			ans = min(ans, len)
			sum -= nums[i]
			i++
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	res := minSubArrayLen(target, nums)
	fmt.Println(res)
}
