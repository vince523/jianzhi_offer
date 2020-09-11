package levelOrderIII


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	// 保存同一层节点
	q := []*TreeNode{root}
	level := 0
	flag := false
	for len(q) != 0 {
		tmp := []*TreeNode{}
		list := make([]int, len(q))
		for i, v := range q {
			// 保存同一层节点的值
			if !flag {
				list[i] = v.Val
			} else {
				list[len(q)-1-i] = v.Val
			}

			// 把左右子节点推到新的队列中
			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}
		}
		ret = append(ret, list)
		// 奇偶层反转
		flag = !flag
		// 层数加一
		level++
		// 重新赋值下一层节点队列
		q = tmp
	}
	return ret
}