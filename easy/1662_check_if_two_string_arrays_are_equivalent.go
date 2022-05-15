package easy

import "bytes"

/*
https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/

Constraints:

1 <= word1.length, word2.length <= 10^3
1 <= word1[i].length, word2[i].length <= 10^3
1 <= sum(word1[i].length), sum(word2[i].length) <= 10^3
word1[i] and word2[i] consist of lowercase letters.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Check If Two String Arrays are Equivalent.
Memory Usage: 2.4 MB, less than 62.07% of Go online submissions for Check If Two String Arrays are Equivalent.
*/

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	b1 := bytes.Buffer{}
	b1.Grow(len(word1))

	for i := 0; i < len(word1); i++ {
		b1.WriteString(word1[i])
	}

	str := b1.String()

	ind := 0
	for i := 0; i < len(word2); i++ {
		for ii := 0; ii < len(word2[i]); ii++ {
			if ind >= len(str) {
				return false
			}
			if str[ind] != word2[i][ii] {
				return false
			}
			ind++

		}
	}
	if ind != len(str) {
		return false
	}
	return true
}
