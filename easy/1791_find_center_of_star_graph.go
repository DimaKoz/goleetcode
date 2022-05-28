package easy

/*
https://leetcode.com/problems/find-center-of-star-graph/

Constraints:

3 <= n <= 105
edges.length == n - 1
edges[i].length == 2
1 <= ui, vi <= n
ui != vi
The given edges represent a valid star graph.
*/

/*
Runtime: 142 ms, faster than 75.00% of Go online submissions for Find Center of Star Graph.
Memory Usage: 17.5 MB, less than 25.00% of Go online submissions for Find Center of Star Graph.
*/

func findCenter(edges [][]int) int {
	edgesCounter := make(map[int]int, len(edges))
	for i := 0; i < len(edges); i++ {
		edgesCounter[edges[i][0]]++
		edgesCounter[edges[i][1]]++
	}
	maxKey := 0
	maxValue := 0
	for key, value := range edgesCounter {
		if value > maxValue {
			maxValue = value
			maxKey = key
		}
	}
	return maxKey
}
