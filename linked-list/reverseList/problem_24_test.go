package reverseList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseList(t *testing.T) {
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

	link := ListNode{
		Val:  0,
		Next: nil,
	}

	// nil
	assert.Nil(t, ReverseList(nil))

	// one node
	assert.Equal(t, &link, ReverseList(&link))

	// reverse
	assert.Equal(t, 4, ReverseList(&link1).Val)
}
