package main

import (
	"context"
	"fmt"
	"runtime"
)

func worker1(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(msg)
		}
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("GOMAXPROCS set to", runtime.GOMAXPROCS(-1))

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	fmt.Scanln()
}
