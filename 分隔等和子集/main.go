package main

import "fmt"

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	// 背包的体积为target
	// 背包要放入的商品重量为元素的数值, 价值也为元素的数值
	dp := make([]int, target+10)
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	if target == dp[target] {
		return true
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	res := canPartition([]int{1, 5, 11, 5})
	fmt.Println(res)
}
