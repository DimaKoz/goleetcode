package easy

/*
https://leetcode.com/problems/check-if-the-sentence-is-pangram/

Constraints:

1 <= sentence.length <= 1000
sentence consists of lowercase English letters.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Check if the Sentence Is Pangram.
Memory Usage: 2.1 MB, less than 65.33% of Go online submissions for Check if the Sentence Is Pangram.
*/

func checkIfPangram(sentence string) bool {
	letters := 26
	if len(sentence) < letters {
		return false
	}
	uniq := 0
	symb := make([]byte, letters)
	ind := 0
	for i := 0; i < len(sentence); i++ {
		ind = int(sentence[i] - 'a')
		if symb[ind] != 0 {
			continue
		}
		symb[ind]++
		uniq++
		if uniq == letters {
			return true
		}
	}
	return false
}
