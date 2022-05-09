package easy

/*
https://leetcode.com/problems/split-a-string-in-balanced-strings/

Constraints:

1 <= s.length <= 1000
s[i] is either 'L' or 'R'.
s is a balanced string.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Split a String in Balanced Strings.
Memory Usage: 1.9 MB, less than 92.11% of Go online submissions for Split a String in Balanced Strings.
*/
func balancedStringSplit(s string) int {
	var rCounter, lCounter, counter int
	lenS := len(s)
	for i := 0; i < lenS; i++ {
		if s[i] == 'R' {
			rCounter++
		} else {
			lCounter++
		}
		if rCounter == lCounter {
			counter++
			rCounter, lCounter = 0, 0
		}
	}
	return counter
}
