package deleteNode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	var	(
		link1, link2, link3, link4 ListNode
	)
	link4 = ListNode{
		Val:  4,
		Next: nil,
	}

	link3 = ListNode{
		Val:  3,
		Next: &link4,
	}

	link2 = ListNode{
		Val:  2,
		Next: &link3,
	}

	link1 = ListNode{
		Val:  1,
		Next: &link2,
	}

	// 删除中间节点
	//DeleteNode(&link1, &link3)
	// 删除尾节点
	DeleteNode(&link1, &link4)
	head := &link1
	var nums []int
	for ;head != nil; head=head.Next{
		nums = append(nums, head.Val)
	}

	// 删除中间节点
	//assert.Equal(t, []int{1, 2, 4}, nums)
	// 删除尾节点
	assert.Equal(t, []int{1, 2, 3}, nums)
}

func TestDeleteDuplication(t *testing.T) {
	var	(
		link1, link2, link3, link4 ListNode
	)
	link4 = ListNode{
		Val:  4,
		Next: nil,
	}

	link3 = ListNode{
		Val:  2,
		Next: &link4,
	}

	link2 = ListNode{
		Val:  2,
		Next: &link3,
	}

	link1 = ListNode{
		Val:  1,
		Next: &link2,
	}

	head := DeleteDuplication(&link1)
	var nums []int
	for ;head!= nil; head=head.Next{
		nums = append(nums, head.Val)
	}
	assert.Equal(t, []int{1, 4}, nums)
}