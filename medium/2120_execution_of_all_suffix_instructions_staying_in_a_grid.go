package medium

/*
https://leetcode.com/problems/execution-of-all-suffix-instructions-staying-in-a-grid/

Constraints:

m == s.length
1 <= n, m <= 500
startPos.length == 2
0 <= start(row,index 0), start(col, index 1) < n
s consists of 'L', 'R', 'U', and 'D'.
*/

/*
Runtime: 11 ms, faster than 100.00% of Go online submissions for Execution of All Suffix Instructions Staying in a Grid.
Memory Usage: 3.9 MB, less than 47.37% of Go online submissions for Execution of All Suffix Instructions Staying in a Grid.
*/

func executeInstructions(n int, startPos []int, s string) []int {
	result := make([]int, len(s))

	var row, col int
	counter := 0
	for i := 0; i < len(s); i++ {
		counter = 0
		row, col = startPos[0], startPos[1]
		for ii := i; ii < len(s); ii++ {
			if s[ii] == 'L' {
				if col > 0 {
					col--
					counter++
				} else {
					break
				}
			} else if s[ii] == 'R' {

				if n > col+1 {
					col++
					counter++
				} else {
					break
				}
			} else if s[ii] == 'U' {
				if row > 0 {
					row--
					counter++
				} else {
					break
				}

			} else if s[ii] == 'D' {
				if n > row+1 {
					row++
					counter++
				} else {
					break
				}
			}
		}
		result[i] = counter
	}
	return result
}
