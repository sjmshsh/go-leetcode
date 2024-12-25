package main

import (
	"fmt"
	"time"
)

// 用channel实现一个限流器

func main() {
	// 每次处理3个请求
	chLimit := make(chan struct{}, 3)
	for i := 0; i < 20; i++ {
		chLimit <- struct{}{}
		go func(i int) {
			fmt.Println("下游服务处理逻辑...", i)
			time.Sleep(time.Second * 3)
			<-chLimit
		}(i)
	}
	time.Sleep(30 * time.Second)
}
