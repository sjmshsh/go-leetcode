package main

import (
	"fmt"
	"time"
)

func PrioritySelect(ch1, ch2 <-chan string) {
	for {
		select {
		case val := <-ch1:
			fmt.Println(val)
		case val2 := <-ch2:
		priority:
			for {
				select {
				case val1 := <-ch1:
					fmt.Println(val1)
				default:
					break priority
				}
			}
			fmt.Println(val2)
		}
	}
}

func main() {
	tick := time.Tick(time.Second)
	for {
		select {
		case t := <-tick:
			fmt.Println(t)
			continue
			fmt.Println("test")
		}
	}
	fmt.Println("end")
}
