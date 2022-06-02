package easy

/*
https://leetcode.com/problems/count-good-triplets/

Constraints:

3 <= arr.length <= 100
0 <= arr[i] <= 1000
0 <= a, b, c <= 1000
*/

/*
Runtime: 9 ms, faster than 77.78% of Go online submissions for Count Good Triplets.
Memory Usage: 2.3 MB, less than 48.15% of Go online submissions for Count Good Triplets.
*/

func countGoodTriplets(arr []int, a int, b int, c int) int {
	counter, resA, resB, resC, i, j, k, lenI, lenK := 0, 0, 0, 0, 0, 0, 0, len(arr)-2, len(arr)-1
	absMinus := func(a int, b int) int {
		result := a - b
		if result < 0 {
			return -result
		}
		return result
	}
	for i = 0; i < lenI; i++ {
		for j = i + 1; j < lenK; j++ {
			for k = j + 1; k < len(arr); k++ {
				resA = absMinus(arr[i], arr[j])
				resB = absMinus(arr[j], arr[k])
				resC = absMinus(arr[i], arr[k])
				if resA <= a && resB <= b && resC <= c {
					counter++
				}
			}
		}
	}
	return counter
}
