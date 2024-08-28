package main

import (
	"context"
	"fmt"
	"time"
)

func f1(in chan struct{}) {
	time.Sleep(time.Second)
	in <- struct{}{}
}

func f2(in chan struct{}) {
	time.Sleep(2 * time.Second)
	in <- struct{}{}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	go func() {
		go f1(ch1)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("f1 timeout")
				break
			case <-ch1:
				fmt.Println("f1 done")
			}
		}
	}()
	go func() {
		go f2(ch2)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("f1 timeout")
				break
			case <-ch2:
				fmt.Println("f2 done")
			}
		}
	}()
	time.Sleep(5 * time.Second)
}
