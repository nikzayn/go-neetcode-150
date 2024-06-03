package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPalindrome(t *testing.T) {
	str := "madam"
	exp := true
	actual := IsPalindrome(str)
	assert.Equal(t, exp, actual)
}
