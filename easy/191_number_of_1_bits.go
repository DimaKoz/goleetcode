package easy

import "math/bits"

/*

https://leetcode.com/problems/number-of-1-bits/

Constraints:

The input must be a binary string of length 32.

*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of 1 Bits.
Memory Usage: 1.8 MB, less than 87.04% of Go online submissions for Number of 1 Bits.
*/
func hammingWeight(num uint32) int {
	var ans = 0
	for num > 0 {
		ans += 1
		num &= num - 1
	}
	return ans
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of 1 Bits.
Memory Usage: 2 MB, less than 40.11% of Go online submissions for Number of 1 Bits.

BenchmarkHammingWeightM   	1000000000	         0.3366 ns/op	       0 B/op	       0 allocs/op
BenchmarkHammingWeight    	 100000000	        11.34 ns/op	           0 B/op	       0 allocs/op
*/
func hammingWeightM(num uint32) int {
	return bits.OnesCount32(num)
}
