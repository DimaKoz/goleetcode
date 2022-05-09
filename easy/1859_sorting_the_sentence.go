package easy

import (
	"bytes"
	"strings"
)

/*
https://leetcode.com/problems/sorting-the-sentence/

Constraints:

2 <= s.length <= 200
s consists of lowercase and uppercase English letters, spaces, and digits from 1 to 9.
The number of words in s is between 1 and 9.
The words in s are separated by a single space.
s contains no leading or trailing spaces.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Sorting the Sentence.
Memory Usage: 2 MB, less than 73.08% of Go online submissions for Sorting the Sentence.
*/
func sortSentence(s string) string {
	spl := strings.Split(s, " ")
	sorted := make([]string, len(spl))
	var word string
	var index int
	for i := 0; i < len(spl); i++ {
		word = spl[i]
		index = int(word[len(word)-1]-48 /*getting a digit from asc II*/) - 1
		sorted[index] = word[:len(word)-1]
	}
	b := &bytes.Buffer{}
	for i := 0; i < len(sorted); i++ {
		if i != 0 {
			b.WriteString(" ")
		}
		b.WriteString(sorted[i])
	}
	return b.String()
}
