package deleteNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteNode(head, node *ListNode)  {
	if head == nil || node == nil {
		return
	}
	if node.Next != nil {
		pNext := node.Next
		node.Val = pNext.Val
		node.Next = pNext.Next
	} else if head == node {
		// 只有一个节点的链表
		node = nil
		head = nil
	} else {
		// 链表存在多个节点，删除尾节点
		for head.Next != node{
			head = head.Next
		}

		head.Next = nil
	}
}
