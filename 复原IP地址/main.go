package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 这是一个切割问题, 但是切割的时候要注意条件判断

var (
	path []string
	res  []string
)

func restoreIpAddresses(s string) []string {
	res = make([]string, 0)
	path = make([]string, 0)
	dfs(s, 0)
	return res
}

func dfs(s string, startIdx int) {
	// 如果已经切分了4份了, 那么就停止
	if len(path) == 4 {
		if startIdx == len(s) {
			str := strings.Join(path, ".")
			res = append(res, str)
		}
		return
	}
	for i := startIdx; i < len(s); i++ {
		// 含有前导0, 无效
		if i != startIdx && s[startIdx] == '0' {
			break
		}
		// 切割字符串
		str := s[startIdx : i+1]
		// 把字符串转换成数字
		num, _ := strconv.Atoi(str)
		if num >= 0 && num <= 255 {
			path = append(path, str)
			dfs(s, i+1)
			path = path[:len(path)-1]
		} else {
			break
		}
	}
}

func main() {
	res := restoreIpAddresses("25525511135")
	fmt.Println(res)
}
