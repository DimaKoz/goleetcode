package easy

import "sort"

/*
https://leetcode.com/problems/maximum-product-difference-between-two-pairs/

Constraints:

4 <= nums.length <= 10^4
1 <= nums[i] <= 10^4
*/

/*
Runtime: 39 ms, faster than 42.86% of Go online submissions for Maximum Product Difference Between Two Pairs.
Memory Usage: 6.6 MB, less than 24.49% of Go online submissions for Maximum Product Difference Between Two Pairs.
*/

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	lN := len(nums)
	return nums[lN-1]*nums[lN-2] - nums[0]*nums[1]
}
