package main

import (
	"fmt"
)

// 本题其实就是尽量让石头分成重量相同的两堆, 相撞之后剩下的石头最小
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	target := sum / 2
	// dp[k]表示容量为j的背包, 最多可以背最大重量为dp[k]
	dp := make([]int, target+10)
	n := len(stones)
	for i := 0; i < n; i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sum - dp[target]*2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	res := lastStoneWeightII([]int{2, 7, 4, 1, 8, 1})
	fmt.Println(res)
}
