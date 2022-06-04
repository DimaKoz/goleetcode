package easy

/*
https://leetcode.com/problems/counting-words-with-a-given-prefix/

Constraints:

1 <= words.length <= 100
1 <= words[i].length, pref.length <= 100
words[i] and pref consist of lowercase English letters.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Counting Words With a Given Prefix.
Memory Usage: 3.5 MB, less than 67.59% of Go online submissions for Counting Words With a Given Prefix.
*/

func prefixCount(words []string, pref string) int {
	counter := 0
	word := ""
	isAllPresent := false
	for i := 0; i < len(words); i++ {
		word = words[i]
		if len(pref) > len(word) {
			continue
		}
		isAllPresent = true
		for ii := 0; ii < len(pref); ii++ {
			if pref[ii] != word[ii] {
				isAllPresent = false
				break
			}
		}
		if isAllPresent {
			counter++
		}
	}
	return counter
}
