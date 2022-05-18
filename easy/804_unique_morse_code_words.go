package easy

import "bytes"

/*
https://leetcode.com/problems/unique-morse-code-words/

Constraints:

1 <= words.length <= 100
1 <= words[i].length <= 12
words[i] consists of lowercase English letters.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Morse Code Words.
Memory Usage: 2.3 MB, less than 74.75% of Go online submissions for Unique Morse Code Words.
*/

func uniqueMorseRepresentations(words []string) int {
	morse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	trans := make(map[string]bool)
	b := bytes.Buffer{}

	for i := 0; i < len(words); i++ {
		b.Reset()
		for ii := 0; ii < len(words[i]); ii++ {
			b.WriteString(morse[words[i][ii]-'a'])
		}
		trans[b.String()] = true
	}
	return len(trans)
}
