package main

import (
	"fmt"
	"time"
)

func main() {
	// 4个channel
	chs := make([]chan int, 4)
	for i, _ := range chs {
		chs[i] = make(chan int)
		// 开启4个协程
		go func(i int) {
			// 从channel当中获取当前通道的值
			v := <-chs[i]
			fmt.Println(v + 1)
			time.Sleep(time.Second)
			// 把下一个值写入到下一个channel
			chs[(i+1)%4] <- (v + 1) % 4
		}(i)
	}
	// 在第一个协程写入0
	chs[0] <- 0
	select {}
}
