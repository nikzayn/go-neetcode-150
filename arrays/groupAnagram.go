package arrays

import "sort"

func GroupAnagrams(str []string) [][]string {
	anagrams := make(map[string][]string, len(str))

	for _, word := range str {
		sortedWord := sortWord(word)
		anagrams[sortedWord] = append(anagrams[sortedWord], word)
	}

	result := [][]string{}
	for _, group := range anagrams {
		result = append(result, group)
	}

	return result
}

func sortWord(word string) string {
	wordBytes := []byte(word)
	sort.Slice(wordBytes, func(i, j int) bool {
		return wordBytes[i] < wordBytes[j]
	})

	return string(wordBytes)
}
