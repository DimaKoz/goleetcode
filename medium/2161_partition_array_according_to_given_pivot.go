package medium

/*
https://leetcode.com/problems/partition-array-according-to-given-pivot/

Constraints:

1 <= nums.length <= 10^5
-10^6 <= nums[i] <= 10^6
pivot equals to an element of nums.
*/

/*
Runtime: 191 ms, faster than 80.95% of Go online submissions for Partition Array According to Given Pivot.
Memory Usage: 9.3 MB, less than 78.57% of Go online submissions for Partition Array According to Given Pivot.
*/

func pivotArray(nums []int, pivot int) []int {
	res := make([]int, len(nums))
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < pivot {
			res[index] = nums[i]
			index++
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == pivot {
			res[index] = nums[i]
			index++
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > pivot {
			res[index] = nums[i]
			index++
		}
	}
	return res
}
