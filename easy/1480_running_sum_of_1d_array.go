package easy

/*

https://leetcode.com/problems/running-sum-of-1d-array/

Constraints:

1 <= nums.length <= 1000
-10^6 <= nums[i] <= 10^6
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Running Sum of 1d Array.
Memory Usage: 2.7 MB, less than 87.84% of Go online submissions for Running Sum of 1d Array.
*/
func runningSum(nums []int) []int {
	lenN := len(nums)
	prev := 0
	for i := 1; i < lenN; i++ {
		nums[i] += nums[prev]
		prev = i
	}
	return nums
}
