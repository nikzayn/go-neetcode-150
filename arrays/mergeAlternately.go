package arrays

/*
Approach - We can try with two pointer - two inputs and exhaust both
  - @Input() -> arr - []int
  - @Output() -> arr - []int
  - TC - O(N + M) | SC - O(N + M)
*/
func MergeAlternately(word1 string, word2 string) string {
	result := ""
	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		result += string(word1[i]) + string(word2[j])
		i++
		j++
	}

	for i < len(word1) {
		result += string(word1[i])
		i++
	}

	for j < len(word2) {
		result += string(word2[j])
		j++
	}

	return result
}
