package main

import "fmt"

// 使用额外的数组
func rotate1(nums []int, k int) {
	newNums := make([]int, k)
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

// 数组翻转
func rotate2(nums []int, k int) {
	// 对k取模
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(a []int) {
	n := len(a)
	left := 0
	right := n - 1
	for left < right {
		// 交换
		a[left], a[right] = a[right], a[left]
		left++
		right--
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate2(nums, k)
	fmt.Println(nums)
}
