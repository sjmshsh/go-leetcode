package main

import "fmt"

func main() {
	s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
	fmt.Println(s)
	quicksort1(s, 0, len(s)-1)
	fmt.Println(s)
}

func quicksort1(s []int, left, right int) {
	if left >= right {
		return
	}
	keyi := partition1(s, left, right)
	quicksort1(s, left, keyi-1) // keyi本身就不用遍历了
	quicksort1(s, keyi+1, right)
}

// 分区，并返回分界点
func partition1(a []int, left, right int) int {
	midIdx := GetMidIndex(a, left, right)
	// 把中间值作为起点，避免最差的时间复杂度
	a[midIdx], a[left] = a[left], a[midIdx]
	keyi := left // 把左边的值作为keyi
	// right找到比keyi位置小的第一个数字
	// left找到比keyi位置大的第一个数字
	for left < right {
		// 右边先走，取小
		for left < right && a[right] >= a[keyi] {
			right--
		}
		for left < right && a[left] <= a[keyi] {
			left++
		}
		a[right], a[left] = a[left], a[right]
	}

	// 把keyi交换到"正确的位置"
	a[keyi], a[left] = a[left], a[keyi]
	return left
}

// 三数取中
func GetMidIndex(a []int, left, right int) int {
	mid := (right + left) / 2
	if a[left] < a[mid] {
		if a[right] < a[left] {
			return left
		} else {
			if a[right] > a[mid] {
				return mid
			} else {
				return right
			}
		}
	} else {
		if a[right] < a[mid] {
			return mid
		} else {
			if a[right] > a[left] {
				return left
			} else {
				return right
			}
		}
	}
}
