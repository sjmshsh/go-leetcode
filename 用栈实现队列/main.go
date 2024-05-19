package main

// MyStack 使用切片实现一个栈
// 由于只是辅助实现算法题目, 因此不做异常情况处理
type MyStack []int

func (s *MyStack) Push(v int) {
	*s = append(*s, v)
}

func (s *MyStack) Pop() int {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func (s *MyStack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *MyStack) Size() int {
	return len(*s)
}

func (s *MyStack) Empty() bool {
	return s.Size() == 0
}

// MyQueue 设置两个栈, 一个输入栈, 一个输出栈
type MyQueue struct {
	stackIn  *MyStack
	stackOut *MyStack
}

func Constructor() MyQueue {
	return MyQueue{
		stackOut: &MyStack{},
		stackIn:  &MyStack{},
	}
}

func (this *MyQueue) Push(x int) {
	this.stackIn.Push(x)
}

func (this *MyQueue) Pop() int {
	this.fillStackOut()
	return this.stackOut.Pop()
}

func (this *MyQueue) Peek() int {
	this.fillStackOut()
	return this.stackOut.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.stackIn.Empty() && this.stackOut.Empty()
}

// 填充输出栈
func (this *MyQueue) fillStackOut() {
	// 如果输出栈为空
	if this.stackOut.Empty() {
		// 就把输入栈里面的所有元素全部插入到输出栈
		for !this.stackIn.Empty() {
			val := this.stackIn.Pop()
			this.stackOut.Push(val)
		}
	}
}
