package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	// 使用一个map来模拟set
	res := []int{}
	set := map[int]struct{}{}
	// 把nums1数组放入到set里面去
	for _, item := range nums1 {
		set[item] = struct{}{}
	}
	for _, item := range nums2 {
		if _, exist := set[item]; exist {
			res = append(res, item)
			// 删除这个元素, 避免重复
			delete(set, item)
		}
	}
	return res
}

func main() {
	nums1 := []int{
		1, 2, 2, 1,
	}
	nums2 := []int{
		2, 2,
	}
	res := intersection(nums1, nums2)
	fmt.Println(res)
}
