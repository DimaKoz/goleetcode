package easy

import "sort"

/*
https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/

Constraints:

2 <= arr.length <= 1000
-10^6 <= arr[i] <= 10^6
*/

/*
Runtime: 4 ms, faster than 52.32% of Go online submissions for Can Make Arithmetic Progression From Sequence.
Memory Usage: 2.4 MB, less than 66.23% of Go online submissions for Can Make Arithmetic Progression From Sequence.
*/

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) == 2 {
		return true
	}
	sort.Ints(arr)
	dif := arr[1] - arr[0]
	for i := 0; i < len(arr)-1; i++ {
		if arr[i]+dif != arr[i+1] {
			return false
		}
	}
	return true
}
