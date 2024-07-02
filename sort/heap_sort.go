package main

import "fmt"

func main() {
	s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
	fmt.Println(s)
	fmt.Println(heapSort(s))
}

func heapSort(s []int) []int {
	// 构建大根堆
	n := len(s)
	for i := (n - 2) / 2; i >= 0; i-- {
		shiftDown(s, i, n)
	}
	// 定义无序边界
	end := n - 1
	for end != 0 {
		s[0], s[end] = s[end], s[0]
		shiftDown(s, 0, end)
		end--
	}
	return s
}

// 大根堆排升序
// 向下调整
func shiftDown(s []int, i, n int) {
	parent := i
	child := i*2 + 1
	// 孩子节点不能越界
	for child < n {
		// 重新定义孩子节点
		if child+1 < n && s[child+1] > s[child] {
			// 取较大的那个孩子节点
			child++
		}
		if s[parent] < s[child] {
			// 交换父亲和孩子的位置
			s[parent], s[child] = s[child], s[parent]
			parent = child
			child = 2*parent + 1
		} else {
			break
		}
	}
}
