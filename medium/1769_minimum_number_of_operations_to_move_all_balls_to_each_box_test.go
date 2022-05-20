package medium

import "testing"

type caseMinOperations struct {
	boxes  string
	expect []int
}

var casesMinOperations = []caseMinOperations{
	{"110", []int{1, 1, 3}},
	{"001011", []int{11, 8, 5, 4, 3, 4}},
}

func TestSlowMinOperationsImpl(t *testing.T) {

	for _, item := range casesMinOperations {
		got := slowMinOperationsImpl(item.boxes)
		for i := 0; i < len(item.expect); i++ {
			if got[i] != item.expect[i] {
				t.Errorf(" Wrong with the case : expect %v, got:%v at:%d", item.expect, got, i)
			}
		}
	}

}

func TestMinOperationsImpl(t *testing.T) {

	for _, item := range casesMinOperations {
		got := minOperations(item.boxes)
		for i := 0; i < len(item.expect); i++ {
			if got[i] != item.expect[i] {
				t.Errorf(" Wrong with the case : expect %v, got:%v at:%d", item.expect, got, i)
			}
		}
	}

}
