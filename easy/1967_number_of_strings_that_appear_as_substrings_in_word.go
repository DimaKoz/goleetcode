package easy

import "strings"

/*
https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/

Constraints:

1 <= patterns.length <= 100
1 <= patterns[i].length <= 100
1 <= word.length <= 100
patterns[i] and word consist of lowercase English letters.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of Strings That Appear as Substrings in Word.
Memory Usage: 2.7 MB, less than 76.92% of Go online submissions for Number of Strings That Appear as Substrings in Word.
*/

func numOfStrings(patterns []string, word string) int {
	counter := 0
	for i := 0; i < len(patterns); i++ {
		if strings.Contains(word, patterns[i]) {
			counter++
		}
	}
	return counter
}
