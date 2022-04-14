package easy

/*
Constraints:

m == mat.length
n == mat[i].length
1 <= m, n <= 100
-1000 <= mat[i][j] <= 1000
1 <= r, c <= 300
*/
func matrixReshape(mat [][]int, r int, c int) [][]int {
	totalItems := len(mat) * len(mat[0])
	if r*c != totalItems {
		return mat
	}
	plain := make([]int, 0, totalItems)
	var row []int
	for _, row = range mat {
		plain = append(plain, row...)
	}
	result := make([][]int, r, r)
	for i := 0; i < r; i++ {
		result[i] = plain[i*c : i*c+c]
	}
	return result
}
