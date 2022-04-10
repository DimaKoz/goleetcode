package easy

import "sort"

func checkStraightLine(coordinates [][]int) bool {

	sort.Slice(coordinates, func(i int, j int) bool {

		for x := range coordinates[i] {
			if coordinates[i][x] == coordinates[j][x] {
				continue
			}
			return coordinates[i][x] < coordinates[j][x]
		}
		return false
	})
	lastIndex := len(coordinates) - 1
	startPoint := coordinates[0]
	endPoint := coordinates[lastIndex]
	for i := 1; i < lastIndex; i++ {
		if !inLine(startPoint, endPoint, coordinates[i]) {
			return false
		}
	}
	return true
}

func inLine(a []int, b []int, p []int) bool {

	if a[0] == p[0] { // if AP is vertical
		return b[0] == p[0]
	}
	if a[1] == p[1] { // if AP is horizontal
		return b[1] == p[1]
	}
	i1 := float32(p[0]-a[0]) / float32(b[0]-a[0])
	i2 := float32(p[1]-a[1]) / float32(b[1]-a[1])

	return i1 == i2 && i1 >= 0.00 && i1 <= 1.00

}
