package arrays

import (
	"slices"
	"strings"
)

/*
Approach - We will approach to solve this problem with inbuilt libraries and two pointer approach exhaust both
  - @Input() -> str - string
  - @Output() -> string
  - TC - O(N) | SC - O(N)
*/
func ReverseWordsLib(s string) string {
	words := strings.Fields(s)
	slices.Reverse(words)
	return strings.Join(words, " ")
}

func ReverseWordsTwoPointer(s string) string {
	words := strings.Fields(s)

	low, high := 0, len(s)-1

	for low < high {
		words[high], words[low] = words[low], words[high]
		low++
		high--
	}

	return strings.Join(words, "")
}
