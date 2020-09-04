package findFirstCommonNode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFirstCommonNode(t *testing.T) {
	var	(
		link1, link2, link3, link4, link5, link6, link7, link8 ListNode
	)
	link4 = ListNode{
		Val:  4,
		Next: &link5,
	}
	link5 = ListNode{
		Val:  5,
		Next: &link6,
	}

	link6 = ListNode{
		Val:  6,
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

	link7 = ListNode{
		Val:  1,
		Next: &link8,
	}

	link8 = ListNode{
		Val:  8,
		Next: &link4,
	}
	result := FindFirstCommonNode(&link1, &link7)
	assert.Equal(t, 4, result.Val)
}