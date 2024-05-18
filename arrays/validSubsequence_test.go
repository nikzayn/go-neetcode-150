package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidSubsequence(t *testing.T) {
	arr := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}
	actual := true
	exp2 := IsValidSubsequenceEff(arr, sequence)
	assert.Equal(t, exp2, actual)
}

func BenchmarkIsValidSubsequence(b *testing.B) {
	arr := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}
	for i := 0; i < b.N; i++ {
		IsValidSubsequenceEff(arr, sequence)
	}
}
