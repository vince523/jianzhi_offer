package buildTree

type TreeNode struct {
 	Val int
 	Left *TreeNode
 	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	for i := range inorder {
		// root 节点
		if inorder[i] == preorder[0] {
			return &TreeNode{
				Val: preorder[0],
				Left: buildTree(preorder[1:i+1], inorder[:]),
				Right: buildTree(preorder[i+1:], inorder[i+1:]),
			}
		}
	}
	return nil
}