package printListReversingly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintListReversingly(t *testing.T) {
	var	(
		link1, link2, link3, link4, link5, link6 ListNode
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

	assert.Equal(t, []int{6, 5, 4, 3, 2, 1}, PrintListReversingly(&link1))
}