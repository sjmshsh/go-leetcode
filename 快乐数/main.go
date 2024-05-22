package main

import "fmt"

func isHappy(n int) bool {
	m := map[int]bool{}
	for n != 1 && !m[n] {
		m[n] = true
		n = getSum(n)
	}
	return n == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		m := n % 10
		n = n / 10
		sum += m * m
	}
	return sum
}

func main() {
	n := 19
	fmt.Println(isHappy(n))
}
