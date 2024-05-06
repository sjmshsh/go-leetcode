package main

import "fmt"

func reverseStr(s string, k int) string {
	arr := []byte(s)
	// 让i每次增加2 * k
	for i := 0; i < len(arr); i += 2 * k {
		if i+k <= len(arr) {
			reverse(arr[i : i+k])
		} else {
			reverse(arr[i:])
		}
	}

	return string(arr)
}

func reverse(str []byte) {
	left := 0
	right := len(str) - 1
	for left < right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
}

func main() {
	var s string
	var k int
	_, err := fmt.Scanf("%s", &s)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Scanf("%d", &k)
	if err != nil {
		panic(err)
	}

	res := reverseStr(s, k)
	fmt.Println(res)
}
