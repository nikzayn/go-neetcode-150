package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidParentheses(t *testing.T) {
	s := "(]"
	exp := false
	actual := ValidParenthees(s)
	assert.Equal(t, exp, actual)
}
