package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidAnagram(t *testing.T) {
	s1 := "leetcode"
	s2 := "codeleet"
	actual := ValidAnagrams(s1, s2)
	exp := true
	assert.Equal(t, exp, actual)
}
