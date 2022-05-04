package easy

/*
https://leetcode.com/problems/shuffle-the-array/

Constraints:

1 <= n <= 500
nums.length == 2n
1 <= nums[i] <= 10^3
*/

/*
Runtime: 4 ms, faster than 81.14% of Go online submissions for Shuffle the Array.
Memory Usage: 3.7 MB, less than 77.71% of Go online submissions for Shuffle the Array.
*/
func shuffle(nums []int, n int) []int {
	lenN := len(nums)
	result := make([]int, lenN)
	counter := 0
	for i := 0; i < lenN; i += 2 {
		result[i] = nums[counter]
		result[i+1] = nums[counter+n]
		counter++
	}
	return result
}
