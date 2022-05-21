package easy

/*
https://leetcode.com/problems/valid-anagram/

Constraints:

1 <= s.length, t.length <= 5 * 10^4
s and t consist of lowercase English letters.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Anagram.
Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for Valid Anagram.
*/

func isAnagram(s string, t string) bool {
	n := len(s)
	if n != len(t) {
		return false
	}
	english := 26
	counts := make([]int, english)

	for i := 0; i < n; i++ {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}
	for i := 0; i < english; i++ {
		if counts[i] != 0 {
			return false
		}
	}

	return true
}
