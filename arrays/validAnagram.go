package arrays

/*
   Observation: First check if length is equal and also how many unicode chars which we need to take care
   Secondly, we need to iterate over both string to store the data in map of rune,

   Then, iterate over rune of map to check if each char is greater than zero, if it is return false, otherwise
   return false

   @Input(): s, t -> string
   @Output(): bool
*/

func ValidAnagrams(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	runeMap := make(map[rune]int)

	for _, v := range s {
		runeMap[v]++
	}

	for _, v := range t {
		runeMap[v]--
	}

	for _, val := range runeMap {
		if val > 0 {
			return false
		}
	}

	return true
}
