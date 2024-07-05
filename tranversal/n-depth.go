package main

import "container/list"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0

	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		n := st.Len()
		for i := 0; i < n; i++ {
			node := st.Remove(st.Front()).(*Node)
			// 把该节点所有的孩子节点都放入队列
			for _, child := range node.Children {
				st.PushBack(child)
			}
		}
		depth++
	}

	return depth
}
