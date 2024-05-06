package main

import "fmt"

func reverseSlice(nums []int) {
	left := 0
	right := len(nums) - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func main() {
	var num int
	var nums []int

	// 读取输入的数字序列
	for {
		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			break
		}
		nums = append(nums, num)
	}

	// 反转切片
	reverseSlice(nums)

	// 输出结果
	fmt.Println(nums)
}
