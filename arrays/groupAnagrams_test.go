package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	exp := [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}
	actual := GroupAnagrams(str)
	assert.Equal(t, exp, actual)
}
