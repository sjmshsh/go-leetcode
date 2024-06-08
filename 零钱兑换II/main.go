package main

import "fmt"

func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 0; i < n; i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

func main() {
	amount := 5
	coins := []int{1, 2, 5}
	res := change(amount, coins)
	fmt.Println(res)
}
