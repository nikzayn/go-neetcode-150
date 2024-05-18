package arrays

/*
Two number sum problem
  - Approach -> In this approach, I will traverse over the array and check
    if the indices values sum up to equal with targetSum or not
  - @Input = arr -> []int, targetSum -> int
  - @Output = []int
  - TC - O(n^2) | SC - O(1)
*/
func TwoNumberSumNaive(arr []int, targetSum int) []int {
	for i := 0; i < len(arr); i++ {
		firstNum := arr[i]
		for j := i + 1; j < len(arr); j++ {
			secondNum := arr[j]
			if firstNum+secondNum == targetSum {
				return []int{firstNum, secondNum}
			}
		}
	}
	return []int{}
}

/*
Two number sum problem
  - Approach -> In this approach, I will traverse over the array and store the pairs in maps
  - @Input = arr -> []int, targetSum -> int
  - @Output = []int
  - TC - O(n) | SC - O(n)
*/
func TwoNumberSumEff(arr []int, targetSum int) []int {
	m := map[int]bool{}
	for _, num := range arr {
		potentialMatch := targetSum - num
		if _, found := m[potentialMatch]; found {
			return []int{potentialMatch, num}
		}
		m[num] = true
	}
	return []int{}
}
