package easy

/*
https://leetcode.com/problems/count-equal-and-divisible-pairs-in-an-array/

Constraints:

1 <= nums.length <= 100
1 <= nums[i], k <= 100
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Count Equal and Divisible Pairs in an Array.
Memory Usage: 2.8 MB, less than 39.19% of Go online submissions for Count Equal and Divisible Pairs in an Array.
*/

func countPairs(nums []int, k int) int {
	counter := 0
	for i := 0; i < len(nums); i++ {
		for ii := i + 1; ii < len(nums); ii++ {
			if nums[i] == nums[ii] {
				if (i*ii)%k == 0 {
					counter++
				}
			}
		}
	}
	return counter
}
