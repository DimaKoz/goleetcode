package easy

import "bytes"

/*
https://leetcode.com/problems/cells-in-a-range-on-an-excel-sheet/

Constraints:

s.length == 5
'A' <= s[0] <= s[3] <= 'Z'
'1' <= s[1] <= s[4] <= '9'
s consists of uppercase English letters, digits and ':'.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Cells in a Range on an Excel Sheet.
Memory Usage: 2.9 MB, less than 93.81% of Go online submissions for Cells in a Range on an Excel Sheet.
*/

func cellsInRange(s string) []string {
	startLetter := s[0]
	endLetter := s[3]
	startDig := s[1]
	endDig := s[4]
	b := bytes.Buffer{}
	b.Grow(2)

	result := make([]string, (endLetter-startLetter+1)*(endDig-startDig+1))
	next := 0
	for i := startLetter; i <= endLetter; i++ {
		for ii := startDig; ii <= endDig; ii++ {
			b.Reset()
			b.WriteByte(i)
			b.WriteByte(ii)
			result[next] = b.String()
			next++
		}
	}
	return result
}
