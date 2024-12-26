package main

import (
	"fmt"
	"time"
)

func print(u <-chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("print int", <-u)
}

func main() {
	c := make(chan int, 5)
	a := 0

	c <- a
	fmt.Println(a)
	a = 1

	go print(c)
	time.Sleep(5 * time.Second)
	fmt.Println(a)
}
