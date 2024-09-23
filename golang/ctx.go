package main

import (
	"context"
	"fmt"
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

}
