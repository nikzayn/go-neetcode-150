package twopointers

func SumOf3Values(nums []int, target int) bool {
	low, high, triplets := 0, 0, 0

	for i := 0; i < len(nums)-2; i++ {
		low = i + 1
		high = len(nums) - 1

		for low < high {
			triplets = nums[i] + nums[low] + nums[high]

			if triplets == target {
				return true
			} else if triplets < target {
				low += 1
			} else {
				high -= 1
			}
		}
	}
	return false
}
