package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructBinaryTree(array []int) *TreeNode {
	var root *TreeNode
	nodes := make([]*TreeNode, len(array))

	// 初始化二叉树节点
	for i := 0; i < len(nodes); i++ {
		var node *TreeNode
		if array[i] != -1 {
			node = &TreeNode{Val: array[i]}
		}
		nodes[i] = node
		if i == 0 {
			root = node
		}
	}
	// 串联节点
	for i := 0; i*2+2 < len(array); i++ {
		if nodes[i] != nil {
			nodes[i].Left = nodes[i*2+1]
			nodes[i].Right = nodes[i*2+2]
		}
	}

	return root
}

func printBinaryTree(root *TreeNode, n int) {
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	}

	result := []int{}
	for len(queue) > 0 {
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			if node != nil {
				result = append(result, node.Val)
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			} else {
				result = append(result, -1)
			}
		}
		// 清除队列中的本层节点, 进入下一层遍历
		queue = queue[len(queue):]
	}

	// 参数n控制输出值数量, 否则二叉树最后一层叶子节点的孩子节点也会被打印(但是这些孩子节点是不存在的).
	fmt.Println(result[:n])
}

func main() {
	array := []int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8}
	root := constructBinaryTree(array)
	printBinaryTree(root, len(array))
}
