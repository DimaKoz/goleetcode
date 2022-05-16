package medium

import (
	"sort"
)

/*
https://leetcode.com/problems/sort-the-matrix-diagonally/

Constraints:

m == mat.length
n == mat[i].length
1 <= m, n <= 100
1 <= mat[i][j] <= 100
*/

/*
Runtime: 3 ms, faster than 100.00% of Go online submissions for Sort the Matrix Diagonally.
Memory Usage: 4.8 MB, less than 92.31% of Go online submissions for Sort the Matrix Diagonally.
*/

func diagonalSort(mat [][]int) [][]int {
	getDiagonalSlice := func(mat [][]int, row int, column int) []int {
		sR := len(mat) - row
		sC := len(mat[0]) - column
		lenD := 0
		if sR > sC {
			lenD = sC
		} else {
			lenD = sR
		}
		res := make([]int, lenD)
		for i := 0; i < len(res); i++ {
			res[i] = mat[row+i][column+i]
		}
		return res
	}
	fillDiagonalSlice := func(m *[][]int, d *[]int, row int, column int) {
		for i := 0; i < len(*d); i++ {
			(*m)[row+i][column+i] = (*d)[i]
		}

	}

	for i := 0; i < len(mat[0])-1; i++ {
		sl := getDiagonalSlice(mat, 0, i)
		sort.Ints(sl)
		fillDiagonalSlice(&mat, &sl, 0, i)
	}
	for i := 1; i < len(mat)-1; i++ {
		sl := getDiagonalSlice(mat, i, 0)
		sort.Ints(sl)
		fillDiagonalSlice(&mat, &sl, i, 0)
	}
	return mat
}
