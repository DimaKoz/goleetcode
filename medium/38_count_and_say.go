package medium

import (
	"bytes"
	"strconv"
)

/*
https://leetcode.com/problems/count-and-say/
Constraints:

1 <= n <= 30
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Count and Say.
Memory Usage: 2.3 MB, less than 96.91% of Go online submissions for Count and Say.
*/

func countAndSay(n int) string {
	s := "1"
	b := &bytes.Buffer{}
	for i := 1; i < n; i++ {
		s = countAndSayInternal(&s, b)
	}
	return s
}

func countAndSayInternal(s *string, b *bytes.Buffer) string {
	b.Reset()
	c := rune((*s)[0])
	count := 1
	lenS := len(*s)
	for i := 1; i < lenS; i++ {
		if rune((*s)[i]) == c {
			count++
		} else {
			b.WriteString(strconv.Itoa(count))
			b.WriteRune(c)
			c = rune((*s)[i])
			count = 1
		}
	}
	b.WriteString(strconv.Itoa(count))
	b.WriteRune(c)
	return b.String()
}
