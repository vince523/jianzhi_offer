package entryNodeOfLoop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntryNodeOfLoop(t *testing.T) {
	var	(
		link1, link2, link3, link4 ListNode
	)
	link4 = ListNode{
		Val:  4,
		Next: &link2,
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

	// 有环
	ret := EntryNodeOfLoop(&link1)
	assert.Equal(t, 2, ret.Val)
	// 没环
	link4.Next = nil
	assert.Nil(t, EntryNodeOfLoop(&link1))
}
