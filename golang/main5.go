package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{
		name: "John",
		age:  30,
	}

	// 获取指向p的指针的反射值，Elem方法用于获取指针指向的值
	v := reflect.ValueOf(&p).Elem()

	// 获取私有字段 name
	nameField := v.FieldByName("name")

	fmt.Println(nameField.String())
}
