package main

import "fmt"

const N = 1e5 + 10

func main() {
	var n int
	fmt.Scan(&n) // 读入数组大小

	w := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&w[i]) // 读取n个整数
	}

	var sum int64 = 0
	for i := 0; i < n; i++ {
		sum += int64(w[i])
	}

	fmt.Println(sum)
}
