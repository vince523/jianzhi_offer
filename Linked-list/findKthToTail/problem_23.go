package problem_03

type ListNode struct {
	Val  int
	Next *ListNode
}

func FindKthToTail(pListHead *ListNode, k int) *ListNode {
	// 空指针判断
	if pListHead == nil || k == 0{
		return nil
	}

	// 双指针
	var p = pListHead
	var q = pListHead
	// 链表索引 用于判断输入的 k 是否会超过总长度
	var i = 0
	for ; p != nil; i++ {
		p = p.Next
		if i > k-1 {
			q = q.Next
		}
	}

	if i > k-1 {
		return q
	}

	return nil
}
