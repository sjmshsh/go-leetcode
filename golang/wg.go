package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	n := 5
	errCh := make(chan error, n) // 用一个缓冲通道来传递错误

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			if i == 3 {
				errCh <- fmt.Errorf("error in goroutine %d", i)
				return
			}
			fmt.Printf("Goroutine %d is done\n", i)
		}(i)
	}

	wg.Wait()
	// 关闭通道以停止接受
	close(errCh)

	for err := range errCh {
		if err != nil {
			fmt.Println("Received Error: ", err)
		}
	}

	fmt.Println("done")
}
