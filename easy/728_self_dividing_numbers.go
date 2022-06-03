package easy

/*
https://leetcode.com/problems/self-dividing-numbers/

Constraints:

1 <= left <= right <= 10^4
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Self Dividing Numbers.
Memory Usage: 2.1 MB, less than 50.00% of Go online submissions for Self Dividing Numbers.
*/

func selfDividingNumbers(left int, right int) []int {
	result := make([]int, 0)
	isAppend := false
	var n, i2 int
	for i := left; i <= right; i++ {
		n = i
		isAppend = true
		for n != 0 {
			i2 = n % 10
			if i2 == 0 {
				isAppend = false
				break
			}
			n /= 10
			if i%i2 != 0 {
				isAppend = false
				break
			}
		}
		if isAppend {
			result = append(result, i)
		}
	}
	return result
}
