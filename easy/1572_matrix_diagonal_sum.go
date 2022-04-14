package easy

/*
Constraints:

n == mat.length == mat[i].length
1 <= n <= 100
1 <= mat[i][j] <= 100

*/

func diagonalSum(mat [][]int) int {
	lenR := len(mat)
	var sum, rtlIndex, i int
	var row []int
	for i, row = range mat {
		sum += row[i]
		rtlIndex = lenR - i - 1
		sum += row[rtlIndex]
	}
	if (lenR & 1) != 0 {
		center := lenR / 2
		sum -= mat[center][center]
	}
	return sum
}
