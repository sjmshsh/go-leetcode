package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type MyType struct {
	A int32
	B float64
}

func MyTypeSliceToBytesSlice(data []MyType) ([]byte, error) {
	buf := new(bytes.Buffer)
	for _, value := range data {
		if err := binary.Write(buf, binary.LittleEndian, value); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func main() {
	// 零切片
	zeroSlice := make([]int, 0)
	fmt.Printf("zeroSlice: %v, len: %d, cap: %d, is nil: %t\n", zeroSlice, len(zeroSlice), cap(zeroSlice), zeroSlice == nil)

	// 空切片
	var emptySlice []int = []int{}
	fmt.Printf("emptySlice: %v, len: %d, cap: %d, is nil: %t\n", emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)

	// nil 切片
	var nilSlice []int
	fmt.Printf("nilSlice: %v, len: %d, cap: %d, is nil: %t\n", nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
}
