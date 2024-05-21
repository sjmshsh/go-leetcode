package main

type MyStack struct {
	// 创建两个队列
	queue1 []int // 主
	queue2 []int // 从
}

func Constructor() MyStack {
	return MyStack{
		queue1: make([]int, 0),
		queue2: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	// 先将数据存在queue2中
	this.queue2 = append(this.queue2, x)
	// 将queue1中所有元素移动到queue2中, 再将两个队列进行交换
	this.Move()
}

func (this *MyStack) Move() {
	if len(this.queue1) == 0 {
		this.queue1, this.queue1 = this.queue2, this.queue1
	} else {
		// queue1元素从头开始一个一个追加到queue2中
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
		this.Move()
	}
}

func (this *MyStack) Pop() int {
	val := this.queue1[0]
	this.queue1 = this.queue2[1:] // 去除第一个元素
	return val
}

func (this *MyStack) Top() int {
	return this.queue1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}

func main() {

}
