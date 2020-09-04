package mergeList

type ListNode struct {
	Val  int
	Next *ListNode
}

func Merge(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var mergeLink *ListNode
	if l1.Val < l2.Val {
		mergeLink = l1
		mergeLink.Next = Merge(l1.Next, l2)
	} else {
		mergeLink = l2
		mergeLink.Next = Merge(l1, l2.Next)
	}

	return mergeLink
}

