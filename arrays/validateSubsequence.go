package arrays

/*
Validate Subsquence
  - Approach -> In this approach, let's check if array consists the subsequence of the master array,
    means the index value should also have same values, it could be any order unrelated to sorted or unsorted
  - @Input = arr -> []int, sequence -> []int
  - @Output = bool
  - TC - O(N^2) | SC - O(1)
*/
// func IsValidSubsequenceNaive(arr []int, sequence []int) bool {
// 	for i := 0; i < len(arr); i++ {
// 		for j := 0; j < len(sequence); j++ {
// 			if arr[i] = sequence[j] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

/*
Validate Subsquence
  - Approach -> In this approach, let's check if array consists the subsequence of the master array,
    means the index value should also have same values, it could be any order unrelated to sorted or unsorted
  - @Input = arr -> []int, sequence -> []int
  - @Output = bool
  - TC - O(N) | SC - O(1)
*/
func IsValidSubsequenceEff(arr []int, sequence []int) bool {
	arrIdx, seqIdx := 0, 0

	for arrIdx < len(arr) && seqIdx < len(sequence) {
		if arr[arrIdx] == sequence[seqIdx] {
			seqIdx += 1
		}
		arrIdx += 1
	}

	return seqIdx == len(sequence)
}
