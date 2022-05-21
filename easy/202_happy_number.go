package easy

/*
https://leetcode.com/problems/happy-number/

Constraints:

1 <= n <= 2^31 - 1
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Happy Number.
Memory Usage: 2 MB, less than 89.20% of Go online submissions for Happy Number.
*/

func isHappy(n int) bool {
	for true {
		n = calcHappyOnce(n)
		if n == 1 {
			return true
		}
		if n < 7 { // 7 - a magick number to do less iterations
			break
		}
	}
	return false
}

func calcHappyOnce(n int) int {
	sum := 0
	for n != 0 {
		d := n % 10
		sum += d * d
		n /= 10
	}
	return sum
}
