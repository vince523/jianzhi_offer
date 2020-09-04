package copyRandomList

type ListNode struct {
	Val    int
	Next   *ListNode
	Random *ListNode
}

func copyRandomList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	newNode := head
	for newNode != nil{
		clone := new(ListNode)
		clone.Next = newNode.Next
		clone.Val = newNode.Val
		clone.Random = nil
		newNode.Next = clone
		newNode = clone.Next
	}

	// 再从头开始
	newNode = head
	for newNode != nil{
		cloneNode := newNode.Next
		if newNode.Random != nil {
			cloneNode.Random=newNode.Random.Next
		}else {
			cloneNode.Random=nil
		}
		newNode = cloneNode.Next
	}

	// 链表拆分, 再从头节点遍历
	newNode = head
	// 复制链表头节点,用于返回值
	cloneHead := head.Next

	// 构造复制的链表
	for newNode != nil {
		cloneNode := newNode.Next
		newNode.Next = cloneNode.Next
		if cloneNode.Next != nil{
			cloneNode.Next = cloneNode.Next.Next
		}

		newNode = newNode.Next
	}
	return cloneHead
}