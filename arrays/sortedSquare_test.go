package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquareNaive(t *testing.T) {
	arr := []int{1, 1, 5, 6, 8, 10, 22, 25}
	actual := []int{1, 1, 25, 36, 64, 100, 484, 625}
	exp := SortedSquareEff(arr)
	assert.Equal(t, exp, actual)
}
