package easy

/*
https://leetcode.com/problems/count-the-number-of-consistent-strings/

Constraints:

1 <= words.length <= 10^4
1 <= allowed.length <= 26
1 <= words[i].length <= 10
The characters in allowed are distinct.
words[i] and allowed contain only lowercase English letters.
*/

/*
Runtime: 39 ms, faster than 68.97% of Go online submissions for Count the Number of Consistent Strings.
Memory Usage: 7.2 MB, less than 64.37% of Go online submissions for Count the Number of Consistent Strings.
*/

func countConsistentStrings(allowed string, words []string) int {
	al := make(map[byte]bool, len(allowed))
	for i := 0; i < len(allowed); i++ {
		al[allowed[i]] = true
	}
	var word string
	var found bool
	counter := 0
	for i := 0; i < len(words); i++ {
		word = words[i]
		for ii := 0; ii < len(word); ii++ {
			if _, found = al[word[ii]]; !found {
				break
			}
		}
		if found {
			counter++
		}
	}
	return counter
}
