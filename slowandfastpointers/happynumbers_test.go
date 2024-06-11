package slowandfastpointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHappy(t *testing.T) {
	num := 23
	exp := true
	actual := isHappy(num)
	assert.Equal(t, exp, actual)
}
