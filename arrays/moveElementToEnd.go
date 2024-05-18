package arrays

/*
Approach - We can try with two pointer approach
  - @Input() -> arr - []int
  - @Output() -> arr - []int
  - TC - O(N) | SC - O(N)
*/
func MoveElementToLast(arr []int, toMove int) []int {
	end := len(arr) - 1
	for start := 0; start <= end; start++ {
		if arr[start] == toMove {
			arr[start], arr[end] = arr[end], arr[start]
			start--
			end--
		}
	}
	return arr
}
