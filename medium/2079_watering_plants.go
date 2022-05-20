package medium

/*
https://leetcode.com/problems/watering-plants/

Constraints:

n == plants.length
1 <= n <= 1000
1 <= plants[i] <= 10^6
max(plants[i]) <= capacity <= 10^9
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Watering Plants.
Memory Usage: 2.8 MB, less than 78.26% of Go online submissions for Watering Plants.
*/

func wateringPlants(plants []int, capacity int) int {
	steps := 0
	units := capacity
	for i := 0; i < len(plants); i++ {
		if units < plants[i] {
			units = capacity
			steps += i * 2
		}
		steps++
		units -= plants[i]
	}
	return steps
}
