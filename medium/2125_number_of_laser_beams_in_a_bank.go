package medium

/*
https://leetcode.com/problems/number-of-laser-beams-in-a-bank/

Constraints:

m == bank.length
n == bank[i].length
1 <= m, n <= 500
bank[i][j] is either '0' or '1'.
*/

/*
Runtime: 22 ms, faster than 95.83% of Go online submissions for Number of Laser Beams in a Bank.
Memory Usage: 8 MB, less than 27.08% of Go online submissions for Number of Laser Beams in a Bank.
*/

func numberOfBeams(bank []string) int {
	prev, current, total := 0, 0, 0
	for i := 0; i < len(bank); i++ {
		if current != 0 {
			prev = current
		}
		current = 0
		for ii := 0; ii < len(bank[i]); ii++ {
			if bank[i][ii] == '1' {
				current++
			}
		}
		total += prev * current
	}
	return total
}
