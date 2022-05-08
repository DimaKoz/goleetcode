package easy

/*
https://leetcode.com/problems/create-target-array-in-the-given-order/

Constraints:

1 <= nums.length, index.length <= 100
nums.length == index.length
0 <= nums[i] <= 100
0 <= index[i] <= i
*/

/*
Runtime: 2 ms, faster than 62.00% of Go online submissions for Create Target Array in the Given Order.
Memory Usage: 2 MB, less than 90.00% of Go online submissions for Create Target Array in the Given Order.
*/
func createTargetArray(nums []int, index []int) []int {
	target := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		target = insertCreateTargetArray(&target, index[i], nums[i])
	}
	return target
}

func insertCreateTargetArray(a *[]int, index int, value int) []int {
	*a = append((*a)[:index+1], (*a)[index:]...)
	(*a)[index] = value
	return *a
}
