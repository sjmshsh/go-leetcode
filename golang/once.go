package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("Resource initialized.")
}

func main() {
	for i := 0; i < 3; i++ {
		go func() {
			once.Do(initialize)
		}()
	}
	fmt.Scanln()
}
