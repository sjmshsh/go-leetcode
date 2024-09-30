package main

import (
	"context"
	"fmt"
	"time"
)

type Rejecthandler func(ctx context.Context) bool

type Queue[T any] interface {
	Push(value T)
	TryPush(value T, timeout time.Duration) bool
	Poll() T
	TryPoll(timeout time.Duration) (T, bool)
}

type BlockingQueue[T any] struct {
	q       chan T // 阻塞队列使用通道表示
	limit   int    // 阻塞队列的大小
	ctx     context.Context
	handler Rejecthandler
}

func NewBlockingQueue[T any](ctx context.Context, queueSize int) *BlockingQueue[T] {
	return &BlockingQueue[T]{
		q:     make(chan T, queueSize),
		limit: queueSize,
		ctx:   ctx,
	}
}

func (q *BlockingQueue[T]) SetRejectHandler(handler Rejecthandler) {
	q.handler = handler
}

func (q *BlockingQueue[T]) Push(value T) {
	ok := true
	if q.handler != nil {
		select {
		case q.q <- value:
			return
		default:
			ok = q.handler(q.ctx)
		}
	}
	if ok {
		q.q <- value
	}
}

func (q *BlockingQueue[T]) TryPush(value T, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(q.ctx, timeout)
	defer cancel()
	select {
	case q.q <- value:
		return true
	case <-ctx.Done():
	}
	return false
}

func (q *BlockingQueue[T]) Poll() T {
	ret := <-q.q
	return ret
}

func (q *BlockingQueue[T]) TryPoll(timeout time.Duration) (ret T, ok bool) {
	ctx, canel := context.WithTimeout(q.ctx, timeout)
	defer canel()
	select {
	case ret = <-q.q:
		return ret, true
	case <-ctx.Done():
	}
	return ret, false
}

func (q *BlockingQueue[T]) size() int {
	return len(q.q)
}

func main() {
	a := NewBlockingQueue[int](context.Background(), 3)
	a.SetRejectHandler(func(ctx context.Context) bool {
		fmt.Println("reject")
		return false
	})
	for i := 0; i < 4; i++ {
		go a.Push(i)
	}
	select {}
}
