package main

import "fmt"

type SingleNode struct {
	Val  int
	Next *SingleNode
}

type MyLinkedList struct {
	dummyHead *SingleNode
	Size      int
}

func Constructor() MyLinkedList {
	newNode := &SingleNode{
		-999,
		nil,
	}
	return MyLinkedList{
		dummyHead: newNode,
		Size:      0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if this == nil || index < 0 || index >= this.Size {
		return -1
	}
	// 让cur等于真正的头节点
	cur := this.dummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &SingleNode{
		Val: val,
	}
	newNode.Next = this.dummyHead.Next // 新节点指向当前头节点
	this.dummyHead.Next = newNode
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &SingleNode{Val: val}
	cur := this.dummyHead
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		index = 0
	} else if index > this.Size { // 如果索引大于链表长度, 直接返回
		return
	}

	newNode := &SingleNode{Val: val}
	cur := this.dummyHead // 设置当前节点为虚拟头节点
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	newNode.Next = cur.Next
	cur.Next = newNode
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size { // 如果索引无效则直接返回
		return
	}
	cur := this.dummyHead        // 设置当前节点为虚拟头节点
	for i := 0; i < index; i++ { // 遍历到要删除节点的前一个节点
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next // 当前节点直接指向下下个节点，即删除了下一个节点
	}
	this.Size-- // 注意删除节点后应将链表大小减一
}

// 打印链表
func (list *MyLinkedList) printLinkedList() {
	cur := list.dummyHead // 设置当前节点为虚拟头节点
	for cur.Next != nil { // 遍历链表
		fmt.Println(cur.Next.Val) // 打印节点值
		cur = cur.Next            // 切换到下一个节点
	}
}
