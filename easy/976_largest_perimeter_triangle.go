package easy

import "sort"

/*
Constraints:

3 <= nums.length <= 10^4
1 <= nums[i] <= 10^6
*/

/*
Runtime: 39 ms, faster than 76.63% of Go online submissions for Largest Perimeter Triangle.
Memory Usage: 7.1 MB, less than 40.76% of Go online submissions for Largest Perimeter Triangle.
*/

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	sum := 0
	largest := 0
	for i := len(nums) - 3; i >= 0; i-- {
		sum = nums[i]+nums[i+1]
		largest = nums[i+2]
		if sum > largest {
			return sum + largest
		}
	}
	return 0
}
