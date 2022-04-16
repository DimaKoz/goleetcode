package easy

/*
Constraints:

1 <= points.length <= 10^4
points[i].length == 2
1 <= x, y, ai, bi <= 10^4
*/

/*
Runtime: 105 ms, faster than 77.34% of Go online submissions for Find Nearest Point That Has the Same X or Y Coordinate.
Memory Usage: 7.3 MB, less than 77.83% of Go online submissions for Find Nearest Point That Has the Same X or Y Coordinate.
*/

// Return the index (0-indexed) of the valid point with the smallest Manhattan distance
// from your current location.
// If there are multiple, return the valid point with the smallest index.
// If there are no valid points, return -1.
func nearestValidPoint(x int, y int, points [][]int) int {
	index := -1
	minManhattanDistance := 10000
	dis := 0
	lenPoints := len(points)
	var point []int

	for i := 0; i < lenPoints; i++ {
		point = points[i]
		if point[0] != x && point[1] != y {
			continue
		}
		if point[0] != x {
			dis = x - point[0]
			if dis < 0 {
				dis = -dis
			}
		} else {
			dis = y - point[1]
			if dis < 0 {
				dis = -dis
			}
		}
		if minManhattanDistance > dis {
			minManhattanDistance = dis
			index = i
		}
	}
	return index
}
