package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseVowels(t *testing.T) {
	str := "leetcode"
	actual := ReverseVowels(str)
	exp := "leotcede"
	assert.Equal(t, exp, actual)
}
