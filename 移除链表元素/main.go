package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	// l是头节点, cur是当前正在遍历的指针
	cur := l
	var res []int
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	var sb strings.Builder
	for _, item := range res {
		sb.WriteString(fmt.Sprintf("%d ", item))
	}
	return sb.String()
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	cur := dummyNode
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyNode.Next
}

func constructList(nums []int) *ListNode {
	head := &ListNode{}
	cur := head
	for _, num := range nums {
		head.Val = num
		head.Next = &ListNode{}
		head = head.Next
	}
	head.Next = nil
	return cur
}

func main() {
	nums := []int{1, 2, 6, 3, 4, 5, 6}
	val := 6
	// 把nums数组转换成链表
	head := constructList(nums)
	node := removeElements(head, val)
	fmt.Println(node)
}
