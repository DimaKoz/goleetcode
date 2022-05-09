package easy

/*
https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/

Constraints:

0 <= num <= 10^6
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of Steps to Reduce a Number to Zero.
Memory Usage: 1.9 MB, less than 14.97% of Go online submissions for Number of Steps to Reduce a Number to Zero.
*/
func numberOfSteps(num int) int {
	counter := 0
	for {
		if num == 0 {
			return counter
		}
		if num&1 == 1 { // odd
			num -= 1
		} else {
			num /= 2
		}
		counter++
	}
}
