package easy

import "testing"

func TestNearestValidPoint(t *testing.T) {

	incomeP := [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}
	incomeX := 3
	incomeY := 4
	expect := 2
	got := nearestValidPoint(incomeX, incomeY, incomeP)
	if got != expect {
		t.Fatal("wrong")
	}
}

func TestNearestValidPointNotValid(t *testing.T) {

	incomeP := [][]int{{2, 3}}
	incomeX := 3
	incomeY := 4
	expect := -1
	got := nearestValidPoint(incomeX, incomeY, incomeP)
	if got != expect {
		t.Fatal("wrong")
	}
}
