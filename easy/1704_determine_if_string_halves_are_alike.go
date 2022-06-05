package easy

/*
https://leetcode.com/problems/determine-if-string-halves-are-alike/

Constraints:

2 <= s.length <= 1000
s.length is even.
s consists of uppercase and lowercase letters.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Determine if String Halves Are Alike.
Memory Usage: 2.2 MB, less than 63.41% of Go online submissions for Determine if String Halves Are Alike.
*/

func halvesAreAlike(s string) bool {
	half := len(s) / 2
	counter1, counter2 := 0, 0
	var l byte
	for i := 0; i < len(s); i++ {
		l = s[i]
		if l == 'a' || l == 'e' || l == 'i' || l == 'o' || l == 'u' || l == 'A' || l == 'E' || l == 'I' || l == 'O' || l == 'U' {
			if i >= half {
				counter2++
			} else {
				counter1++
			}
		}
	}
	return counter2 == counter1
}
