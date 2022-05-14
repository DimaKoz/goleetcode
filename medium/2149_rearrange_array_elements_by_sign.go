package medium

/*
https://leetcode.com/problems/rearrange-array-elements-by-sign/

Constraints:

2 <= nums.length <= 2 * 10^5
nums.length is even
1 <= |nums[i]| <= 10^5
nums consists of equal number of positive and negative integers.
*/

/*
Runtime: 269 ms, faster than 62.22% of Go online submissions for Rearrange Array Elements by Sign.
Memory Usage: 13.6 MB, less than 44.44% of Go online submissions for Rearrange Array Elements by Sign.
*/

func rearrangeArray(nums []int) []int {
	lenS := len(nums)
	result := make([]int, lenS)
	indP := 0
	indN := 1
	item := 0
	for i := 0; i < lenS; i++ {
		item = nums[i]
		if item > 0 {
			result[indP] = item
			indP += 2
		} else {
			result[indN] = item
			indN += 2
		}
	}
	return result
}
