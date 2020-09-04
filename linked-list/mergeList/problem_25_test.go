package mergeList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	var	(
		link1, link2, link3, link4, link5, link6 ListNode
	)
	link4 = ListNode{
		Val:  4,
		Next: &link6,
	}
	link5 = ListNode{
		Val:  5,
		Next: nil,
	}

	link6 = ListNode{
		Val:  6,
		Next: nil,
	}

	link3 = ListNode{
		Val:  3,
		Next: &link5,
	}

	link2 = ListNode{
		Val:  2,
		Next: &link4,
	}

	link1 = ListNode{
		Val:  1,
		Next: &link3,
	}

	link := Merge(&link1, &link2)
	nums := make([]int, 0)
	for link != nil {
		nums = append(nums, link.Val)
		link = link.Next
	}

	//normal case
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, nums)

}