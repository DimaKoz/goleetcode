package easy

/*
https://leetcode.com/problems/number-of-good-pairs/

Constraints:

1 <= nums.length <= 100
1 <= nums[i] <= 100
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of Good Pairs.
Memory Usage: 2 MB, less than 87.50% of Go online submissions for Number of Good Pairs.
*/

func numIdenticalPairs(nums []int) int {
	counter := 0
	for i := 0; i < len(nums); i++ {
		for ii := i + 1; ii < len(nums); ii++ {
			if nums[i] == nums[ii] {
				counter++
			}
		}
	}
	return counter
}
