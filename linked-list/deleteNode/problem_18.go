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

func DeleteDuplication(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	temp := &ListNode{0, nil}
	result := temp

	flag := false
	for ;head != nil; head=head.Next {
		pNext := head.Next
		// 相等
		if pNext != nil && pNext.Val == head.Val {
			flag = true
		}else {
			// 不相等情况
			// 先判断flag
			if flag {
				flag = false
			} else {
				// 加入 temp
				temp.Next = &ListNode{head.Val, nil}
				temp = temp.Next
			}
		}
	}

	return result.Next
}