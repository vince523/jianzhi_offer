package kthLargest

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	nums := make([]int, 0)
	reverse(root, &nums)
	return nums[k-1]
}

func reverse(root *TreeNode, nums *[]int) {
	if root.Right != nil {
		reverse(root.Right, nums)
	}
	if root != nil {
		*nums = append(*nums, root.Val)
	}
	if root.Left != nil {
		reverse(root.Left, nums)
	}
}
