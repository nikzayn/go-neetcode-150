package arrays

import "sort"

/*
Approach - We have an array of values which denotes coins value and we need to check how many changes that we cannot create from the

	** values available
	 - @Input() -> arr - []int
	 - @Output() -> arr - []int
	 - TC - O(nlogn) | SC - O(N)
*/
func NonConstructibleChange(arr []int) int {
	sort.Ints(arr)
	currentChange := 0

	for _, val := range arr {
		if val > currentChange+1 {
			return currentChange + 1
		}
		currentChange += val
	}

	return currentChange + 1
}
