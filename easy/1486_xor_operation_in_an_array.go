package easy

/*
https://leetcode.com/problems/xor-operation-in-an-array/

Constraints:

1 <= n <= 1000
0 <= start <= 1000
n == nums.length
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for XOR Operation in an Array.
Memory Usage: 1.8 MB, less than 100.00% of Go online submissions for XOR Operation in an Array.
*/
func xorOperation(n int, start int) int {
	result, counter := start, start
	for i := 0; i < n; i++ {
		if start != counter {
			result ^= counter
		}
		counter += 2
	}
	return result
}
