package arrays

/*
Approach - We will approach to solve this problem with two pointer approach
  - @Input() -> s, t - string
  - @Output() -> bool
  - TC - O(N) | SC - O(1)
*/
func IsSubsequence(s, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}
