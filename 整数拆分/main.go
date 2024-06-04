package main

import "fmt"

func integerBreak(n int) int {
	// 动态规划五部曲
	// 1. 确定dp下标及其含义
	// 2. 确定递推公式
	// 3. 确定dp初始化
	// 4. 确定遍历顺序
	// 5. 打印dp
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 2; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), dp[i-j]*j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	res := integerBreak(10)
	fmt.Println(res)
}
