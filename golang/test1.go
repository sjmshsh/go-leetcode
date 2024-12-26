package main

import (
	"fmt"
	"time"
)

type people struct {
	name string
}

var u = people{name: "A"}

func printPeople(u <-chan *people) {
	time.Sleep(2 * time.Second)
	fmt.Println("printPeople", <-u)
}

func main() {
	c := make(chan *people, 5)
	var a = &u

	c <- a
	fmt.Println(a)
	a = &people{name: "B"}

	go printPeople(c)
	time.Sleep(5 * time.Second)
	fmt.Println(a)
}
