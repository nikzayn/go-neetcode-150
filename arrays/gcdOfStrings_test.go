package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGcdOfStrings(t *testing.T) {
	s1, s2 := "ABABAB", "ABAB"
	exp := "AB"
	actual := GcdOfStrings(s1, s2)
	assert.Equal(t, exp, actual)
}
