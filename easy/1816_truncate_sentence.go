package easy

/*
https://leetcode.com/problems/truncate-sentence/

Constraints:

1 <= s.length <= 500
k is in the range [1, the number of words in s].
s consist of only lowercase and uppercase English letters and spaces.
The words in s are separated by a single space.
There are no leading or trailing spaces.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Truncate Sentence.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for Truncate Sentence.
*/

func truncateSentence(s string, k int) string {
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			k--
		}
		if k == 0 {
			return s[:i]
		}
	}
	return s
}
