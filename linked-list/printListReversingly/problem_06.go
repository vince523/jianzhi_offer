package printListReversingly

import "container/list"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintListReversingly(head *ListNode) []int {
	st := list.New()
	for head != nil {
		st.PushFront(head.Val)
		head = head.Next
	}

	var nums []int
	for eles := st.Front(); eles != nil; eles=eles.Next() {
		nums = append(nums, eles.Value.(int))
	}
	return nums
}