package main

import (
	"fmt"
	"math"
)

// 思路:
// 既然为target, 那么就一定有left组合 - right组合 = target
// left + right = sum, 而 sum 是固定的, 即有right = sum - left
// 因此可以推导出left = ( target + sum ) / 2
// 因此问题就是在集合nums当中找到和为left的组合
func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if abs(target) > sum {
		return 0
	}

	if (target+sum)%2 != 0 {
		return 0
	}

	bags := (target + sum) / 2

	// 在nums当中找到和为bags的组合
	dp := make([]int, bags+10)
	dp[0] = 1
	// dp[j]的含义是当背包的容量为j的时候, 有dp[j]种不同的方案
	for i := 0; i < n; i++ {
		for j := bags; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bags]
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	res := findTargetSumWays(nums, target)
	fmt.Println(res)
}
