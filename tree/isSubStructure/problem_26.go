package isSubStructure

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}
	var res bool
	if A.Val == B.Val {
		res = DoesTreeAHasTreeB(A, B)
	}
	if !res {
		res = isSubStructure(A.Left, B)
	}
	if !res {
		res = isSubStructure(A.Right, B)
	}
	return res
}

func DoesTreeAHasTreeB(A, B *TreeNode) bool {
	// 说明越界匹配完成
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}

	if A.Val != B.Val {
		return false
	}

	return DoesTreeAHasTreeB(A.Left, B.Left) && DoesTreeAHasTreeB(A.Right, B.Right)
}