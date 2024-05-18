package arrays

import "sort"

/*
Approach - In this approach, we are thinking to use an empty slice in which we store the squared index value and
later than we can sort it out
  - @Input() -> arr - []int
  - @Output() -> arr - []int
  - TC - O(nlogn) | SC - O(N)
*/
func SortedSquareNaive(arr []int) []int {
	sortedArr := make([]int, len(arr))

	for i, val := range arr {
		sortedArr[i] = val * val
	}

	sort.Ints(sortedArr)
	return sortedArr
}

/*
Approach - In this approach, we are going with two pointer approach to provide the sorted square
  - @Input() -> arr - []int
  - @Output() -> arr - []int
  - TC - O(n) | SC - O(N)
*/
func SortedSquareEff(arr []int) []int {
	sortedArr := make([]int, len(arr))
	start := 0
	end := len(arr) - 1

	for i := len(arr) - 1; i >= 0; i-- {
		smallVal := arr[start]
		largestVal := arr[end]

		if abs(smallVal) > abs(largestVal) {
			sortedArr[i] = smallVal * smallVal
			start += 1
		} else {
			sortedArr[i] = largestVal * largestVal
			end -= 1
		}
	}
	return sortedArr
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
