package main

import (
	"fmt"
	"sort"
)

// 直接排序
func sortedSquares1(nums []int) []int {
	ans := make([]int, len(nums))
	for i, v := range nums {
		ans[i] = v * v
	}
	sort.Ints(ans)
	return ans
}

// 根据题目升序排序的特点解题
func sortedSquares2(nums []int) []int {
	n := len(nums)
	// 找出正负的分界线
	negative := 0
	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			negative = i
		} else {
			break
		}
	}
	res := make([]int, 0, n)
	// 以分界线
	i := negative
	j := negative + 1
	for i >= 0 || j < n {
		if i < 0 {
			res = append(res, nums[j]*nums[j])
			j++
		} else if j == n {
			res = append(res, nums[i]*nums[i])
			i--
		} else if nums[i]*nums[i] < nums[j]*nums[j] {
			res = append(res, nums[i]*nums[i])
			i--
		} else {
			res = append(res, nums[j]*nums[j])
			j++
		}
	}
	return res
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	res := sortedSquares2(nums)
	fmt.Println(res)
}
