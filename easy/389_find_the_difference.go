package easy


/*
Constraints:

0 <= s.length <= 1000
t.length == s.length + 1
s and t consist of lowercase English letters.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Find the Difference.
Memory Usage: 2.1 MB, less than 94.51% of Go online submissions for Find the Difference.
*/
func findTheDifference(s string, t string) byte {
	symbols := make(map[uint8]int)
	ls := len(s)
	lt := ls + 1
	found := false
	var value, i int
	var key uint8
	for i = 0; i < ls; i++ {
		key = s[i]
		if value, found = symbols[key]; !found {
			symbols[key] = 1
		} else {
			symbols[key] = value + 1
		}
	}

	for i = 0; i < lt; i++ {
		key = t[i]
		if value, found = symbols[key]; found {
			value-=1
			if value == 0 {
				delete(symbols, key)
			} else {
				symbols[key] = value
			}

		} else {
			return key
		}
	}
	return 0
}
