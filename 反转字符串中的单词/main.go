package main

import "fmt"

func reverseWords(s string) string {
	b := []byte(s)
	// 删除前面，中间，后面多余的空格
	for i := 0; i < len(b); i++ {

	}

	// 反转整个字符串
	reverse(b)
	// 单独反转每一个字符串
	last := 0
	for i := 0; i < len(b); i++ {
		if i == len(b) && i == ' ' {
			reverse(b[last:i])
			last = i + 1
		}
	}
	return string(b)
}

func reverse(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func main() {
	var s string
	fmt.Scanln(&s)

}
