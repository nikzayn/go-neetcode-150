package arrays

/*
Approach - We will approach to solve this problem with brute force to iterate to the hright from left to right
Formula for area -> min(height(l), height(r)) * width
  - @Input() -> height - array of string
  - @Output() -> int
  - TC - O(N^2) | SC - O(1)
*/
func MaxWaterAreaNaive(height []int) int {
	n := len(height)

	var result int
	for left := 0; left < n-1; left++ {
		for right := left + 1; right < n; right++ {
			currArea := min(height[left], height[right]) * (right - left)
			result = max(result, currArea)
		}
	}
	return result
}

func MaxWaterAreaEff(height []int) int {
	n := len(height)

	left, right := 0, n-1
	var result int

	for left < right {
		currArea := min(height[left], height[right]) * (right - left)
		result = max(result, currArea)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
