package easy

/*
Constraints:

1 <= nums.length <= 10^4
-2^31 <= nums[i] <= 2^31 - 1
*/

func moveZeroes(nums []int) {
	var next, temp int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			//swap
			temp = nums[next]
			nums[next] = nums[i]
			nums[i] = temp

			next++
		}
	}
}
