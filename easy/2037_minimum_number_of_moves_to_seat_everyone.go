package easy

import "sort"

/*
https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/

Constraints:

n == seats.length == students.length
1 <= n <= 100
1 <= seats[i], students[j] <= 100
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Minimum Number of Moves to Seat Everyone.
Memory Usage: 3.4 MB, less than 65.31% of Go online submissions for Minimum Number of Moves to Seat Everyone.
*/

func minMovesToSeat(seats []int, students []int) int {
	moves := 0
	sort.Ints(seats)
	sort.Ints(students)

	diff := 0
	for i := 0; i < len(students); i++ {
		diff = students[i] - seats[i]
		if diff < 0 {
			diff = -diff
		}
		moves += diff
	}
	return moves
}
