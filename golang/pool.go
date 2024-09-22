package main

import "sync"

var pool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1024)
	},
}

func process() {
	b := pool.Get().([]byte)
	// 使用b进行操作
	// 归还到对象池
	pool.Put(b)
}
