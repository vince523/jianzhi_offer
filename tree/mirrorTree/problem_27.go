package mirrorTree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right  = mirrorTree(root.Right), mirrorTree(root.Left)

	return root
}