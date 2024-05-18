package arrays

/*
Observation:

		Create a map to store the nums data
		Iterate over the nums slice and the value in map
		Create a bucket of [][]int, iterate over map and store the append the data to bucket slice
		Create a slice with k as a capacity
		Iterate over bucket slice in descending order as values are already sorted.
	    Then, append the value to ans slice from bucket
*/
func TopKElements(nums []int, k int) []int {
	m := make(map[int]int)

	for _, val := range nums {
		m[val]++
	}

	bucket := make([][]int, len(nums)+1)

	for key, val := range m {
		bucket[val] = append(bucket[val], key)
	}

	ans := make([]int, 0, k)

	for i := len(bucket) - 1; i >= 0; i-- {
		for _, val := range bucket[i] {
			if k > 0 {
				ans = append(ans, val)
				k--
			}
		}
	}

	return ans
}
