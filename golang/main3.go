package main

import (
	"fmt"
	"time"
)

var (
	word = make(chan struct{})
	num  = make(chan struct{})
)

func printNum() {
	for i := 0; i < 10; i++ {
		<-word
		fmt.Println(i)
		num <- struct{}{}
	}
}

func printWord() {
	for i := 0; i < 10; i++ {
		<-num
		fmt.Println("a")
		word <- struct{}{}
	}
}

func main() {
	go printNum()
	go printWord()
	num <- struct{}{}
	time.Sleep(time.Second * 1)
}
