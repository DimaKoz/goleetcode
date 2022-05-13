package easy

/*
https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/

Constraints:

1 <= nums.length <= 200
1 <= nums[i] <= 100
1 <= k <= 99

*/

/*
Runtime: 4 ms, faster than 94.67% of Go online submissions for Count Number of Pairs With Absolute Difference K.
Memory Usage: 4.1 MB, less than 16.00% of Go online submissions for Count Number of Pairs With Absolute Difference K.
*/
func countKDifference(nums []int, k int) int {
	var result = 0

	counts := make(map[int]int)
	var value int
	var found bool
	for i := 0; i < len(nums); i++ {
		value, found = counts[nums[i]]
		counts[nums[i]] = value + 1
	}
	for i := 0; i < len(nums); i++ {
		if value, found = counts[nums[i]-k]; found {
			result += value
		}
	}
	return result
}
