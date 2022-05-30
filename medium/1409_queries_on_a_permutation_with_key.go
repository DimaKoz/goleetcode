package medium

/*
https://leetcode.com/problems/queries-on-a-permutation-with-key/

Constraints:

1 <= m <= 10^3
1 <= queries.length <= m
1 <= queries[i] <= m
*/

/*
Runtime: 5 ms, faster than 60.87% of Go online submissions for Queries on a Permutation With Key.
Memory Usage: 9.2 MB, less than 8.70% of Go online submissions for Queries on a Permutation With Key.
*/

func processQueries(queries []int, m int) []int {
	moveElementToFirstIndex := func(slice []int, s int) []int {
		x := append([]int{slice[s]}, slice[:s]...)
		x = append(x, slice[s+1:]...)
		return x
	}
	findPos := func(x []int, y int) int {
		for i := 0; i < len(x); i++ {
			if y == x[i] {
				return i
			}
		}
		return -1
	}
	p := make([]int, m)
	for i := 0; i < m; i++ {
		p[i] = i + 1
	}
	for i := 0; i < len(queries); i++ {
		position := findPos(p, queries[i])
		p = moveElementToFirstIndex(p, position)
		queries[i] = position
	}
	return queries
}
