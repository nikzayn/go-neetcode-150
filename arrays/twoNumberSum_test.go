package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoNumberSum(t *testing.T) {
	arr := []int{2, 7, 11, 15}
	target := 9
	actual := []int{2, 7}
	exp := TwoNumberSumEff(arr, target)
	assert.Equal(t, exp, actual)
}
