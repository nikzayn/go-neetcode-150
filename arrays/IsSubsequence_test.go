package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	s1, s2 := "abc", "ahbgdc"
	exp := true
	actual := IsSubsequence(s1, s2)
	assert.Equal(t, exp, actual)
}
