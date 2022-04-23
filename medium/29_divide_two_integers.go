package medium

/*
https://leetcode.com/problems/divide-two-integers/

Constraints:

-2^31 <= dividend, divisor <= 2^31 - 1
divisor != 0
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Divide Two Integers.
Memory Usage: 2.2 MB, less than 89.54% of Go online submissions for Divide Two Integers.
*/

func divide(dividend int, divisor int) int {
	const MaxInt32 = 2147483647
	result := int64(dividend) / int64(divisor)
	if result > MaxInt32 {
		return MaxInt32
	}
	return int(result)
}
