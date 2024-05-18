package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKElements(t *testing.T) {
	nums := []int{1, 1, 1, 1, 2, 2, 3}
	k := 2
	exp := []int{1, 2}
	actual := TopKElements(nums, k)
	assert.Equal(t, exp, actual)
}
