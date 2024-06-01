package main

import "fmt"

// 思路：
// 先使用回溯算法求出所有的子串
// 然后在回溯的过程当中判断是否是回文传即可

var (
	res  [][]string
	path []string
)

func partition(s string) [][]string {
	path = make([]string, 0)
	res = make([][]string, 0)
	dfs(s, 0)
	return res
}

func dfs(s string, startIdx int) {
	// 如果起始位置等于s的大小, 说明已经找到了一组分隔方案了
	if startIdx == len(s) {
		tmp := make([]string, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := startIdx; i < len(s); i++ {
		str := s[startIdx : i+1]
		if isPalindrome(str) {
			path = append(path, str)
			dfs(s, i+1)
			path = path[:len(path)-1]
		}
	}
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	res := partition("aab")
	fmt.Println(res)
}
