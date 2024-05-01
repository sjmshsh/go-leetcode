package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	for in.Scan() {
		str := in.Text()
		s := strings.Split(str, " ")
		sort.Strings(s) // 排序
		// 把切片连接成字符串
		fmt.Println(strings.Join(s, " "))
	}
}
