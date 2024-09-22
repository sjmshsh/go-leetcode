package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	if err := trace.Start(f); err != nil {
		fmt.Println(err)
		return
	}
	defer trace.Stop()

	for i := 0; i < 1000; i++ {
		go func(i int) {
			fmt.Println("Hello World, ", i)
		}(i)
	}
}
