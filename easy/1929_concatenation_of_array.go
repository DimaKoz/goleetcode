package easy

/*

https://leetcode.com/problems/concatenation-of-array/

Constraints:

n == nums.length
1 <= n <= 1000
1 <= nums[i] <= 1000
*/

/*
Runtime: 8 ms, faster than 84.97% of Go online submissions for Concatenation of Array.
Memory Usage: 6.2 MB, less than 70.62% of Go online submissions for Concatenation of Array.
*/
func getConcatenation(nums []int) []int {
	nums = append(nums, nums...)
	return nums
}

/*
Runtime: 19 ms, faster than 11.16% of Go online submissions for Concatenation of Array.
Memory Usage: 6.3 MB, less than 70.62% of Go online submissions for Concatenation of Array.
*/
func getConcatenationV1(nums []int) []int {
	lenN := len(nums)
	result := make([]int, lenN*2)
	ii := lenN
	for i := 0; i < lenN; i++ {
		result[i], result[ii] = nums[i], nums[i]
		ii++
	}
	return result

}
