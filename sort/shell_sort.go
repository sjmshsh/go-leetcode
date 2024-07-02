package main

import "fmt"

func main() {
	s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
	fmt.Println(s)
	fmt.Println(shellSort(s))
}

func shellSort(s []int) []int {
	// 不断缩小gap的值
	for gap := len(s) / 2; gap > 0; gap = gap / 2 {
		// 进行一个直接插入排序, 遍历每一个被gap分开的区间
		for i := gap; i < len(s); i++ {
			for j := i; j >= gap && s[j-gap] > s[j]; j = j - gap {
				s[j], s[j-gap] = s[j-gap], s[j]
			}
		}
	}
	return s
}
