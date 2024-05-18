package stacks

/*
   Observation - Use two stacks to solve this problem, why because we can store the brackets accordingly to that open and
   closed scenario and also we can map of rune to store brackets
   Use two slices
   Input() - s -> string
   Output() - bool
*/

func ValidParenthees(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stacks := []rune{}

	for _, char := range s {
		if _, ok := pairs[char]; ok {
			stacks = append(stacks, char)
		} else if len(stacks) == 0 || pairs[stacks[len(stacks)-1]] != char {
			return false
		} else {
			stacks = stacks[:len(stacks)-1]
		}
	}

	return len(stacks) == 0
}
