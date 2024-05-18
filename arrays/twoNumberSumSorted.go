package arrays

/*

 */

func TwoNumberSumSorted(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	{
		switch {
		case nums[l]+nums[r] < target:
			r--
		case nums[r]+nums[l] > target:
			l++
		default:
			return []int{l + 1, r + 1}
		}
	}
}
