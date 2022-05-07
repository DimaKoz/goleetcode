package easy

/*
https://leetcode.com/problems/decode-xored-array/

Constraints:

2 <= n <= 10^4
encoded.length == n - 1
0 <= encoded[i] <= 10^5
0 <= first <= 10^5
*/

/*
Runtime: 32 ms, faster than 76.32% of Go online submissions for Decode XORed Array.
Memory Usage: 7.5 MB, less than 31.58% of Go online submissions for Decode XORed Array.
*/
func decode(encoded []int, first int) []int {
	result := make([]int, len(encoded)+1)
	result[0] = first
	i1 := 0
	for i := 0; i < len(encoded); i++ {
		i1 = i + 1
		result[i1] = first ^ encoded[i]
		first = result[i1]
	}
	return result
}
