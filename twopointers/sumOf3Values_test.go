package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOf3Values(t *testing.T) {
	s := []int{3, 7, 1, 2, 8, 4, 5}
	target := 10
	exp := true
	actual := SumOf3Values(s, target)
	assert.Equal(t, exp, actual)
}
