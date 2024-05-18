package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeAlternately(t *testing.T) {
	word1, word2 := "abcd", "pq"
	exp := "apbqcd"
	actual := MergeAlternately(word1, word2)
	assert.Equal(t, exp, actual)
}
