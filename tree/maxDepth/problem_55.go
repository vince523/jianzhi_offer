package maxDepth

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var level int
	q := []*TreeNode{root}
	for len(q) != 0 {
		tmp := make([]*TreeNode, 0)
		for i := 0; i < len(q); i++ {
			node := q[i]
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		level++
		q = tmp
	}
	return level
}