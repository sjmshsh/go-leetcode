package main

import "fmt"

var (
	m    []string
	path []byte
	res  []string
)

func letterCombinations(digits string) []string {
	m = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path = make([]byte, 0, len(digits))
	res = make([]string, 0)
	if digits == "" {
		return res
	}
	dfs(digits, 0)
	return res
}

func dfs(digits string, start int) {
	if len(path) == len(digits) {
		tmp := string(path)
		res = append(res, tmp)
		return
	}
	digit := int(digits[start] - '0')
	str := m[digit-2] // 取数字对应的字符集
	for j := 0; j < len(str); j++ {
		path = append(path, str[j])
		dfs(digits, start+1)
		path = path[:len(path)-1]
	}
}

func main() {
	res := letterCombinations("23")
	fmt.Println(res)
}
