package easy

/*
https://leetcode.com/problems/majority-element/
Constraints:

n == nums.length
1 <= n <= 5 * 10^4
-10^9 <= nums[i] <= 10^9
*/

/*
Runtime: 16 ms, faster than 87.44% of Go online submissions for Majority Element.
Memory Usage: 6.1 MB, less than 89.24% of Go online submissions for Majority Element.
*/

func majorityElement(nums []int) int {
	m := make(map[int]int) // map[value of nums[i]]counter
	maxCounter := 0
	valueForMaxCounter := 0
	found := false
	temp := 0
	halfL := len(nums) / 2
	item := 0
	for _, item = range nums {
		if temp, found = m[item]; found {
			temp++
		} else {
			temp = 1
		}
		if temp > maxCounter {
			maxCounter = temp
			valueForMaxCounter = item
			if halfL < maxCounter {
				break
			}
		}
		m[item] = temp
	}
	return valueForMaxCounter
}
