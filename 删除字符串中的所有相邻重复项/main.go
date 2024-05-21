package main

import "fmt"

func removeDuplicates(s string) string {
	// 准备一个栈
	st := []rune{}
	for _, c := range s {
		// 判断栈顶的元素和当前元素是否是相同的
		if len(st) > 0 && st[len(st)-1] == c {
			st = st[:len(st)-1]
		} else {
			st = append(st, c)
		}
	}
	return string(st)
}

func main() {
	s := "abbaca"
	duplicates := removeDuplicates(s)
	fmt.Println(duplicates)
}
