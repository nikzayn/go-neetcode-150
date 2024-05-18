package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxWater(t *testing.T) {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	exp := 49
	actual := MaxWaterAreaEff(arr)
	assert.Equal(t, exp, actual)
}
