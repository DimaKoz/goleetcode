package easy

import (
	"sort"
)

/*
 Constraints:

 1 <= arr.length <= 500
 0 <= arr[i] <= 104
*/

func sortByBits(arr []int) []int {
	var iBits, jBits int8
	sort.Slice(arr, func(i int, j int) bool {
		// or we can do bits.OnesCount instead fn
		fn := func(x int) int8 {
			var c int8 = 0
			for c = 0; x > 0; c++ {
				x &= x - 1 // clear the least significant bit set
			}
			return c
		}
		iBits = fn(arr[i])
		jBits = fn(arr[j])
		if iBits == jBits {
			return arr[i] < arr[j]
		}
		return iBits < jBits
	})
	return arr
}
