package easy

import (
	"bytes"
)

/*
Constraints:

1 <= word1.length, word2.length <= 100
word1 and word2 consist of lowercase English letters.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Strings Alternately.
Memory Usage: 2 MB, less than 99.24% of Go online submissions for Merge Strings Alternately.
*/
func mergeAlternately(word1 string, word2 string) string {
	lenW1 := len(word1)
	lenW2 := len(word2)
	var b bytes.Buffer //use bytes instead of rune cose - only lowercase English letters
	var i int
	for i = 0; i < lenW1 && i < lenW2; i++ {
		b.WriteByte(word1[i])
		b.WriteByte(word2[i])
	}
	if lenW2 > lenW1 {
		b.WriteString(word2[i:])
	} else if lenW2 < lenW1 {
		b.WriteString(word1[i:])
	}
	return b.String()
}
