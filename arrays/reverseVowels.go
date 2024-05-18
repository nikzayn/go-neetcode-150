package arrays

import "strings"

/*
Approach - We will approach to solve this problem with two pointer approach
later than we can sort it out
  - @Input() -> str - string
  - @Output() -> str - string
  - TC - O(N) | SC - O(!)
*/
func ReverseVowels(str string) string {
	low, high := 0, len(str)-1
	vowels := "aeiouAEIOU"
	runeStr := []rune(str)

	for low < high {
		for low < high && !strings.Contains(vowels, string(runeStr[low])) {
			low++
		}

		for low < high && !strings.Contains(vowels, string(runeStr[high])) {
			high--
		}

		if low < high {
			runeStr[low], runeStr[high] = runeStr[high], runeStr[low]
			low++
			high--
		}
	}
	return string(runeStr)
}
