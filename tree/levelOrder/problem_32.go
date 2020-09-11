package levelOrder

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func levelOrder(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; i < len(q); i++ {
		node := q[i]
		ret = append(ret, node.Val)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return ret
}