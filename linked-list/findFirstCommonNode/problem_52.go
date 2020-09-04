package findFirstCommonNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func FindFirstCommonNode(headA *ListNode, headB *ListNode) *ListNode {
	if headA == nil || headB == nil{
		return nil
	}

	lenA := getListNodeLength(headA)
	lenB := getListNodeLength(headB)
	var diff int
	if lenA > lenB {
		for diff = lenA - lenB; diff > 0; diff=diff-1 {
			headA = headA.Next
		}
	} else {
		for diff = lenB - lenA; diff > 0; diff=diff-1 {
			headB = headB.Next
		}
	}

	for ;headA != nil && headB != nil; headA, headB=headA.Next, headB.Next{
		if headA == headB {
			return headA
		}
	}
	return nil
}

func getListNodeLength(node *ListNode) int {
	var count = 0
	for ;node != nil; node=node.Next {
		count++
	}
	return count
}