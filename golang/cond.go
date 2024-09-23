package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	for i := 0; i < 10; i++ {
		go func(i int) {
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait() // Wait方法会将调用的goroutine放入等待队列, 并解锁与Cond关联的互斥锁
			fmt.Printf("Goroutine %d is awake\n", i)
		}(i)
	}

	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		cond.Signal() // 被唤醒的goroutine会依次打印消息, 唤醒等待队列中的一个goroutine. 如果等待队列为空则什么也不会发生
		// 它通常在特定的条件达成后被调用, 以唤醒一个正在等待继续执行的goroutine
		time.Sleep(100 * time.Millisecond)
	}
}
