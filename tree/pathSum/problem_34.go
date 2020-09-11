package pathSum

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	recur(root, sum, []int{}, &ret)
	return ret
}

func recur(root *TreeNode, sum int, arr []int, ret *[][]int) {
	if root == nil {
		return
	}

	arr = append(arr, root.Val)
	// 遍历到最后一层
	if root.Val == sum && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*ret = append(*ret, tmp)
	}

	recur(root.Left, sum-root.Val, arr, ret)
	recur(root.Right, sum-root.Val, arr, ret)
	arr = arr[:len(arr)-1]
}
