package medium

/*
https://leetcode.com/problems/subrectangle-queries/
Constraints:

There will be at most 500 operations considering both methods: updateSubrectangle and getValue.
1 <= rows, cols <= 100
rows == rectangle.length
cols == rectangle[i].length
0 <= row1 <= row2 < rows
0 <= col1 <= col2 < cols
1 <= newValue, rectangle[i][j] <= 10^9
0 <= row < rows
0 <= col < cols
*/

/*
Runtime: 86 ms, faster than 20.34% of Go online submissions for Subrectangle Queries.
Memory Usage: 9.2 MB, less than 8.47% of Go online submissions for Subrectangle Queries.
*/
type SubrectangleQueries struct {
	rec [][]int
}

/*
it was renamed from Constructor to ConstructorSubrectangleQueries
*/
func ConstructorSubrectangleQueries(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	lenR := len(this.rec)
	lenC := 0
	var row []int
	for i := 0; i < lenR; i++ {
		if i < row1 {
			continue
		}
		if i > row2 {
			continue
		}
		row = this.rec[i]
		lenC = len(row)
		for ii := 0; ii < lenC; ii++ {
			if ii < col1 {
				continue
			}
			if ii > col2 {
				continue
			}
			row[ii] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.rec[row][col]
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
