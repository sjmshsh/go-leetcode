package main

import (
	"fmt"
	"time"
)

func main() {
	// 4个channel
	// 这里初始化的是一个数组，元素是4个，元素类型是chan int
	chs := make([]chan int, 4)
	for i, _ := range chs {
		chs[i] = make(chan int)
		// 开4个协程
		go func(i int) {
			for {
				// 获取当前channel值并打印
				v := <-chs[i]
				fmt.Println(v + 1)
				time.Sleep(time.Second)
				// 把下一个值写入下一个channel，等待下一次消费
				chs[(i+1)%4] <- (v + 1) % 4
			}
		}(i)
	}
	// 往第一个塞入0
	chs[0] <- 0
	select {}
}
