package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 1000; i++ {
		writer.WriteString("Hello World")
	}

	writer.Flush()
}
