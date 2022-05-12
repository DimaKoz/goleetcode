package easy

/*
https://leetcode.com/problems/count-of-matches-in-tournament/

Constraints:

1 <= n <= 200
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Count of Matches in Tournament.
Memory Usage: 1.9 MB, less than 10.17% of Go online submissions for Count of Matches in Tournament.
*/

func numberOfMatches(n int) int {
	matches := 0
	for {
		if n <= 1 {
			return matches
		}
		if n&1 == 1 { //odd
			matches += (n - 1) / 2
			n = (n-1)/2 + 1
		} else {
			matches += n / 2
			n /= 2
		}
	}
}
