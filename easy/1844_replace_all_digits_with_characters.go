package easy

import "bytes"

/*
https://leetcode.com/problems/replace-all-digits-with-characters/

Constraints:

1 <= s.length <= 100
s consists only of lowercase English letters and digits.
shift(s[i-1], s[i]) <= 'z' for all odd indices i.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Replace All Digits with Characters.
Memory Usage: 2 MB, less than 85.71% of Go online submissions for Replace All Digits with Characters.
*/

func replaceDigits(s string) string {
	b := bytes.Buffer{}
	b.Grow(len(s))
	var letter, shift, afterShift byte
	for i := 0; i < len(s); i += 2 {
		letter = s[i]
		b.WriteByte(letter)
		if i+1 < len(s) {
			shift = s[i+1] - '0'
			afterShift = letter + shift
			b.WriteByte(afterShift)
		}
	}
	return b.String()
}
