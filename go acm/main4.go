package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := bufio.NewScanner(os.Stdin)
	for inputs.Scan() {
		data := strings.Split(inputs.Text(), " ")
		var sum int
		for i := range data {
			val, _ := strconv.Atoi(data[i]) // 把字符串转换成int
			sum += val
		}
		fmt.Println(sum)
	}
}
