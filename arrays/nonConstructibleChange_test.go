package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonConstructibleChange(t *testing.T) {
	arr := []int{1, 2, 5}
	exp := 4
	actual := NonConstructibleChange(arr)
	assert.Equal(t, actual, exp)
}
