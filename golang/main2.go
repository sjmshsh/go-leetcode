package main

import (
	"fmt"
	"time"
)

var (
	dog  = make(chan struct{})
	cat  = make(chan struct{})
	fish = make(chan struct{})
)

func Dog() {
	<-fish
	fmt.Println("dog")
	dog <- struct{}{}
}

func Cat() {
	<-dog
	fmt.Println("cat")
	cat <- struct{}{}
}

func Fish() {
	<-cat
	fmt.Println("fish")
	fish <- struct{}{}
}

func main() {
	for i := 0; i < 100; i++ {
		go Dog()
		go Cat()
		go Fish()
	}
	fish <- struct{}{}
	time.Sleep(10 * time.Second)
}
