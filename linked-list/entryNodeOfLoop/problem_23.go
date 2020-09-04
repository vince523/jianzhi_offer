package problem_23

type ListNode struct {
	Val int
	Next *ListNode
}

func EntryNodeOfLoop(head *ListNode)  *ListNode{
	if head == nil {
		return nil
	}

	var (
		quick = head
		slow = head
	)

	for {
		// 没有环
		if quick == nil || quick.Next == nil{
			return nil
		}

		if quick.Next == quick {
			return quick
		}

		quick = quick.Next.Next
		slow = slow.Next
		if quick == slow {
			break
		}
	}

	// quick == slow 不一定是环的节点
	// 因此，要再从头走一波链表，两者都是以步长为1走，直到相遇
	quick = head
	for quick != slow {
		quick = quick.Next
		slow = slow.Next
	}
	return quick

}