package easy

/*
https://leetcode.com/problems/decompress-run-length-encoded-list/

Constraints:

2 <= nums.length <= 100
nums.length % 2 == 0
1 <= nums[i] <= 100
*/

/*
Runtime: 3 ms, faster than 82.76% of Go online submissions for Decompress Run-Length Encoded List.
Memory Usage: 6.1 MB, less than 78.16% of Go online submissions for Decompress Run-Length Encoded List.
*/

func decompressRLElist(nums []int) []int {
	result := make([]int, 0, len(nums)) //try to anticipate)
	var freq, val, i, ii int
	for i = 0; i < len(nums); i += 2 {
		freq = nums[i]
		val = nums[i+1]
		for ii = 0; ii < freq; ii++ {
			result = append(result, val)
		}
	}
	return result
}
