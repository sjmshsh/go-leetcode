package main

import "container/list"

func postorderTraversal(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return nil
	}

	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		ans = append(ans, node.Val)

		if node.Left != nil {
			st.PushBack(node.Left)
		}
		if node.Right != nil {
			st.PushBack(node.Right)
		}
	}

	reverse(ans)
	return ans
}

func reverse(a []int) {
	l := 0
	r := len(a) - 1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}
