package easy

/*
https://leetcode.com/problems/two-sum/
Constraints:

2 <= nums.length <= 10^4
-10^9 <= nums[i] <= 10^9
-10^9 <= target <= 10^9
Only one valid answer exists.
*/

/*
Runtime: 3 ms, faster than 97.02% of Go online submissions for Two Sum.
Memory Usage: 4.3 MB, less than 24.84% of Go online submissions for Two Sum.
*/

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	lenNums := len(nums)
	item := 0
	var value int
	var found bool
	for i := 0; i < lenNums; i++ {
		item = nums[i]
		second := target - item
		if value, found = numsMap[second]; found {
			return []int{i, value}
		}
		numsMap[item] = i
	}
	return []int{}
}
