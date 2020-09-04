package problem_03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthToTail(t *testing.T) {
	link3 := ListNode{
		Val:  3,
		Next: nil,
	}

	link2 := ListNode{
		Val:  2,
		Next: &link3,
	}

	link1 := ListNode{
		Val:  1,
		Next: &link2,
	}

	ret := FindKthToTail(&link1, 2)
	assert.Equal(t, ret.Val, 2)
	ret = FindKthToTail(&link1, 4)
	assert.Equal(t, nil, nil)
	ret = FindKthToTail(&link1, 0)
	assert.Equal(t, nil, nil)
}
