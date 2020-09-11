package isBalanced

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return math.Abs(float64(height(root.Left))-float64(height(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}


func height(root *TreeNode) int {
	if root == nil {
		// 后面+1 了
		return -1
	}
	// 树高度是左右子树较高者+1
	return max(height(root.Left), height(root.Right))+1
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}
