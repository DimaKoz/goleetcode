package easy

/*
Constraints:

1 <= words.length <= 100
1 <= words[i].length <= 20
order.length == 26
All characters in words[i] and order are English lowercase letters.

*/

/*
A solution offered by leetcode has been changed a bit.
Runtime: 0 ms, faster than 100.00% of Go online submissions for Verifying an Alien Dictionary.
Memory Usage: 2.5 MB, less than 96.62% of Go online submissions for Verifying an Alien Dictionary.
*/

func isAlienSorted(words []string, order string) bool {
	lenOrder := len(order)
	orderMap := make(map[uint8]uint8, lenOrder)

	var i, i1, j, lenWordI, lenWordI1 int
	for i = 0; i < lenOrder; i++ {
		orderMap[order[i]] = uint8(i)
	}
	var wordI, wordI1 string
	var s, s1 uint8
	for i = 0; i < len(words)-1; i++ {
		i1 = i + 1
		wordI = words[i]
		wordI1 = words[i1]
		lenWordI = len(wordI)
		lenWordI1 = len(wordI1)
		// If we do not find a mismatch letter between words[i] and words[i + 1],
		// we need to examine the case when words are like ("apple", "app").
		for j = 0; j < lenWordI; j++ {
			if j >= lenWordI1 {
				return false
			}
			s = wordI[j]
			s1 = wordI1[j]
			if s != s1 {
				if orderMap[s] > orderMap[s1] {
					return false
				}
				// if we find the first different letter and they are sorted,
				// then there's no need to check remaining letters
				break
			}
		}
	}
	return true
}
