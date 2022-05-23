package easy

import "sort"

/*
https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/

Constraints:

2 <= nums.length <= 500
0 <= nums[i] <= 100
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for How Many Numbers Are Smaller Than the Current Number.
Memory Usage: 3.8 MB, less than 8.88% of Go online submissions for How Many Numbers Are Smaller Than the Current Number.
*/

func smallerNumbersThanCurrent(nums []int) []int {
	cpy := make([]int, len(nums))
	copy(cpy, nums)
	sort.Ints(cpy)

	m := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if _, found := m[cpy[i]]; !found {
			m[cpy[i]] = i
		}
	}
	for i := 0; i < len(nums); i++ {
		cpy[i] = m[nums[i]]
	}
	return cpy
}
