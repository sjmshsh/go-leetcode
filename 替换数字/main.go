package main

import "fmt"

func main() {
	var b []byte
	fmt.Scanln(&b)
	for i := 0; i < len(b); i++ {
		if b[i] >= '0' && b[i] <= '9' {
			// 如果这个字符是数字的话
			increaseElem := []byte{'n', 'u', 'm', 'b', 'e', 'r'}
			b = append(b[:i], append(increaseElem, b[i+1:]...)...)
			i = i + len(increaseElem) - 1
		}
	}
	fmt.Println(string(b))
}
