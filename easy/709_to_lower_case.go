package easy

import (
	"bytes"
)

/*
Constraints:

1 <= s.length <= 100
s consists of printable ASCII characters.
*/
func toLowerCase(s string) string {
	// by the task we can't use strings.ToLower(s string)
	lenS := len(s)
	var b bytes.Buffer
	b.Grow(lenS)
	var c uint8
	const diffASCII = 'a' - 'A'
	for i := 0; i < lenS; i++ {
		c = s[i]
		if 'A' <= c && c <= 'Z' {
			c += diffASCII
		}
		b.WriteByte(c)
	}
	return b.String()
}
