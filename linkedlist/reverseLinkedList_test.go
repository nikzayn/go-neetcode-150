package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TC - O(1) | SC - (1)
func TestReverseLinkedListIterative(t *testing.T) {
	s := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	exp := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}
	actual := ReverseLinkedListIterative(s)
	assert.Equal(t, exp, actual)
}
