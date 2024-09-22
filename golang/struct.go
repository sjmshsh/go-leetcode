package main

import (
	"encoding/json"
	"fmt"
)

type SampleStruct struct {
	Name string
	Data []int // 切片类型不可比较
}

func main() {
	s1 := SampleStruct{Name: "example", Data: []int{1, 2, 3}}
	s2 := SampleStruct{Name: "example", Data: []int{1, 2, 3}}

	b1, _ := json.Marshal(s1)
	b2, _ := json.Marshal(s2)

	fmt.Println(string(b1) == string(b2)) // true
}
