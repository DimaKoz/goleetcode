package medium

import "testing"

type caseProcessQueries struct {
	queries []int
	m       int
	expect  []int
}

var casesProcessQueries = []caseProcessQueries{
	{[]int{3, 1, 2, 1}, 5, []int{2, 1, 2, 1}},
	{[]int{4, 1, 2, 2}, 4, []int{3, 1, 2, 0}},
	{[]int{7, 5, 5, 8, 3}, 8, []int{6, 5, 0, 7, 5}},
}

func TestProcessQueries(t *testing.T) {

	for _, item := range casesProcessQueries {
		got := processQueries(item.queries, item.m)
		for i := 0; i < len(item.expect); i++ {
			if got[i] != item.expect[i] {
				t.Errorf(" Wrong with the case : expect %v, got:%v at:%d", item.expect, got, i)
			}
		}
	}

}
