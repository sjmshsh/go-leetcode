package main

import "fmt"

func findDuplicate(nums []int) int {
	n := len(nums)
	for i := 0; i < n; {
		x := nums[i]   // 这个数字
		index := x - 1 // 这个数字的下标, 在正确的位置上
		// 检查当前数字是不是在正确的位置上面
		if nums[index] == x {
			// 在正确的位置上
			if index != i {
				return x
			}
			i++
		} else {
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	res := findDuplicate(nums)
	fmt.Println(res)
}
