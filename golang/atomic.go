package main

import (
	"fmt"
	"sync"
)

func main() {
	var pool = sync.Pool{
		New: func() interface{} {
			return new(string)
		},
	}

	// 获取对象
	obj := pool.Get().(*string)
	fmt.Println("Object from pool: ", *obj)

	// 重置对象
	*obj = "Hello, sync.Pool"

	// 放回池中
	pool.Put(obj)

	// 再次获取对象
	obj2 := pool.Get().(*string)
	fmt.Println("Object from pool again: ", *obj2)
}
