package easy

import "sort"

/*Constraints:

1 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9
*/

func containsDuplicate(nums []int) bool {
	return containsDuplicate1(nums)
}

/*
Runtime: 96 ms, faster than 52.85% of Go online submissions for Contains Duplicate.
Memory Usage: 7.8 MB, less than 96.81% of Go online submissions for Contains Duplicate.
*/
func containsDuplicate2(nums []int) bool { // less memory
	numsLen := len(nums)
	if numsLen == 1 {
		return false
	}
	sort.Ints(nums)
	for i := 1; i < numsLen; i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}

/*
Runtime: 66 ms, faster than 85.17% of Go online submissions for Contains Duplicate.
Memory Usage: 9.3 MB, less than 30.31% of Go online submissions for Contains Duplicate.
*/
func containsDuplicate1(nums []int) bool { // more fast
	numsLen := len(nums)
	if numsLen == 1 {
		return false
	}
	found := false
	already := make(map[int]bool)
	key := 0
	for i := 0; i < numsLen; i++ {
		key = nums[i]
		if _, found = already[key]; found {
			return true
		}
		already[key] = true
	}
	return false

}
