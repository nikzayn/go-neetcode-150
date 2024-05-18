package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveElementToEnd(t *testing.T) {
	arr := []int{2, 1, 2, 2, 2, 3, 4, 2}
	toMove := 2
	actual := MoveElementToLast(arr, toMove)
	exp := []int{4, 1, 3, 2, 2, 2, 2, 2}
	assert.Equal(t, exp, actual)
}
