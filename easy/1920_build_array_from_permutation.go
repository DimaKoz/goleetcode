package easy

/*
https://leetcode.com/problems/build-array-from-permutation/

Constraints:

1 <= nums.length <= 1000
0 <= nums[i] < nums.length
The elements in nums are distinct.
*/

/*
Runtime: 17 ms, faster than 54.22% of Go online submissions for Build Array from Permutation.
Memory Usage: 6.6 MB, less than 49.10% of Go online submissions for Build Array from Permutation.
*/

func buildArray(nums []int) []int {
	lenN := len(nums)
	result := make([]int, lenN)
	for i := 0; i < lenN; i++ {
		result[i] = nums[nums[i]]
	}
	return result
}
