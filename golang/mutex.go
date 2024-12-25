package main

import (
	"log"
	"time"
)

// 用channel实现一个互斥锁

type Mutex struct {
	ch chan struct{}
}

// NewMutex 初始化锁
func NewMutex() *Mutex {
	mu := &Mutex{
		ch: make(chan struct{}, 1),
	}
	mu.ch <- struct{}{}
	return mu
}

// Lock 加锁，阻塞获取
func (m *Mutex) Lock() {
	<-m.ch // 让ch里面为空即加锁
}

// Unlock 释放锁
func (m *Mutex) Unlock() {
	select {
	// 成功写入channel代表释放锁成功
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

// TryLock 尝试获取锁
func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}

func (m *Mutex) LockTimeout(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)

	select {
	case <-m.ch:
		// 成功获取关闭定时器
		timer.Stop()
		return true
	case <-timer.C:
	}

	// 获取锁超时
	return false
}

// IsLocked 是否上锁
func (m *Mutex) IsLocked() bool {
	return len(m.ch) == 0
}

func main() {
	m := NewMutex()
	ok := m.TryLock()
	log.Printf("locked v %v\n", ok)
	ok = m.TryLock()
	log.Printf("locked v %v\n", ok)

	go func() {
		time.Sleep(5 * time.Second)
		m.Unlock()
	}()

	ok = m.LockTimeout(10 * time.Second)
	log.Printf("LockTimeout v %v\n", ok)
}
