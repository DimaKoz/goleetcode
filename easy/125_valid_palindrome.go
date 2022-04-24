package easy

import "bytes"

/*
https://leetcode.com/problems/valid-palindrome/
Constraints:

1 <= s.length <= 2 * 10^5
s consists only of printable ASCII characters.
*/

/*
Runtime: 3 ms, faster than 73.37% of Go online submissions for Valid Palindrome.
Memory Usage: 2.8 MB, less than 71.68% of Go online submissions for Valid Palindrome.
*/
func isPalindrome(s string) bool {
	b := &bytes.Buffer{}
	lenS := len(s)
	b.Grow(lenS)

	var c uint8
	const diffASCII = 'a' - 'A' // to lower case
	for i := 0; i < lenS; i++ {
		c = s[i]
		if 'A' <= c && c <= 'Z' { // to lower case
			c += diffASCII
		} else if !('a' <= c && c <= 'z' || '0' <= c && c <= '9') { // out of alphanumeric range
			continue
		}
		b.WriteByte(c)
	}

	candidate := b.String()
	lenC := len(candidate)
	if lenC == 0 {
		return true
	}
	for i := 0; i < lenC; i++ {
		if candidate[i] != candidate[lenC-i-1] {
			return false
		}
	}

	return true

}
