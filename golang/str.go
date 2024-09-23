package main

import (
	"fmt"
)

func main() {
	// 原字符串
	str := "我爱你"

	// 将字符串转换为 rune 数组(适用于处理Unicode字符)
	runeArray := []rune(str)

	runeArray[1] = 'a'

	modifiedStr := string(runeArray)
	fmt.Println(modifiedStr)
}
