package easy

/*
https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/

Constraints:

1 <= sentences.length <= 100
1 <= sentences[i].length <= 100
sentences[i] consists only of lowercase English letters and ' ' only.
sentences[i] does not have leading or trailing spaces.
All the words in sentences[i] are separated by a single space.
*/

/*
Runtime: 2 ms, faster than 87.20% of Go online submissions for Maximum Number of Words Found in Sentences.
Memory Usage: 3.6 MB, less than 66.35% of Go online submissions for Maximum Number of Words Found in Sentences.
*/
func mostWordsFound(sentences []string) int {
	max := 1
	var current int
	lenS := len(sentences)
	phrase := ""
	lenP := 0
	for i := 0; i < lenS; i++ {
		current = 1
		phrase = sentences[i]
		lenP = len(phrase)
		for ii := 0; ii < lenP; ii++ {
			if phrase[ii] == ' ' {
				current++
				if current > max {
					max++
				}
			}
		}
	}
	return max
}
