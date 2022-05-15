package easy

/*
https://leetcode.com/problems/rings-and-rods/

Constraints:

rings.length == 2 * n
1 <= n <= 100
rings[i] where i is even is either 'R', 'G', or 'B' (0-indexed).
rings[i] where i is odd is a digit from '0' to '9' (0-indexed).
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Rings and Rods.
Memory Usage: 2 MB, less than 67.53% of Go online submissions for Rings and Rods.
*/

func countPoints(rings string) int {
	r := make([]int, 10)
	g := make([]int, 10)
	b := make([]int, 10)
	counted := make([]int, 10)
	counter := 0
	rod := 0
	var color byte
	for i := 0; i < len(rings); i += 2 {
		rod = int(rings[i+1] - '0')
		if counted[rod] > 0 {
			continue
		}
		color = rings[i]
		if color == 'R' {
			r[rod]++
		} else if color == 'G' {
			g[rod]++
		} else {
			b[rod]++
		}
		if r[rod] > 0 && g[rod] > 0 && b[rod] > 0 {
			counter++
			counted[rod]++
		}
	}
	return counter
}
