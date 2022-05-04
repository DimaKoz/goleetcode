package easy

/*
https://leetcode.com/problems/jewels-and-stones/

Constraints:

1 <= jewels.length, stones.length <= 50
jewels and stones consist of only English letters.
All the characters of jewels are unique.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Jewels and Stones.
Memory Usage: 2.1 MB, less than 52.11% of Go online submissions for Jewels and Stones.
*/

func numJewelsInStones(jewels string, stones string) int {
	m := make(map[byte]bool, 0)
	counter := 0
	var i int
	var found bool
	for i = 0; i < len(jewels); i++ {
		m[jewels[i]] = true
	}
	for i = 0; i < len(stones); i++ {
		if _, found = m[stones[i]]; found {
			counter++
		}
	}
	return counter
}
