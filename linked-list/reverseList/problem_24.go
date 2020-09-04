package reverseList

type ListNode struct {
	Val int
	Next *ListNode
 }

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 初始化头节点的前一个节点
	var prev *ListNode = nil

	for head != nil {
		prev, head, head.Next = head, head.Next, prev
	}

	return prev
}
