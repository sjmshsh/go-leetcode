package main

import "fmt"

// 快慢指针法
func removeElement(nums []int, val int) int {
	// 初始化慢指针
	slow := 0
	fast := 0
	n := len(nums)
	for fast < n {
		for fast < n && nums[fast] == val {
			fast++
		}
		if fast == n {
			break
		}
		// 每一轮都需要
		nums[slow] = nums[fast]
		slow++
		fast++
	}
	return slow
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	res := removeElement(nums, val)
	fmt.Println(res)
}
