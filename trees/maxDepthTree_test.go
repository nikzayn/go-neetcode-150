package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TC - O(N) | SC - (1)
func TestMaxDepthTree(t *testing.T) {
	s := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	exp := 3
	actual := MaxDepth(s)
	assert.Equal(t, exp, actual)
}
