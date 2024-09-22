package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from: ", r)
		}
	}()
	fmt.Println("Start main")
	panic("A server error occurred!")
	fmt.Println("End main")
}

func deferExample() {
	for i := 0; i < 3; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}
