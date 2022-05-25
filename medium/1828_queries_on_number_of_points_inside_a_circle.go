package medium

/*
https://leetcode.com/problems/queries-on-number-of-points-inside-a-circle/

Constraints:

1 <= points.length <= 500
points[i].length == 2
0 <= xi, yi <= 500
1 <= queries.length <= 500
queries[j].length == 3
0 <= xj, yj <= 500
1 <= rj <= 500
All coordinates are integers.
*/

/*
Runtime: 35 ms, faster than 63.16% of Go online submissions for Queries on Number of Points Inside a Circle.
Memory Usage: 6.6 MB, less than 52.63% of Go online submissions for Queries on Number of Points Inside a Circle.
*/

func countPoints(points [][]int, queries [][]int) []int {

	result := make([]int, len(queries))
	var rPow, x3, y3, distancePow2 int
	for i := 0; i < len(queries); i++ {
		rPow = queries[i][2] * queries[i][2]
		for ii := 0; ii < len(points); ii++ {
			// a distance from a center of the circle == √((x1−h)^2+(y1−k)^2)
			// a distance^2 from a center of the circle == x3+y3
			// which we have to compare with r^2
			// where x3 = (x1−h)^2 , points[ii][0] == x1, h = queries[i][0]
			x3 = points[ii][0] - queries[i][0]
			x3 *= x3
			y3 = points[ii][1] - queries[i][1]
			y3 *= y3
			distancePow2 = x3 + y3

			if distancePow2 <= rPow {
				result[i]++
			}

		}
	}
	return result
}
