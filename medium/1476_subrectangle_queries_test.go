package medium

import "testing"

func TestSubrectangleQueries(t *testing.T) {

	srq := ConstructorSubrectangleQueries([][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}})

	if srq.GetValue(0, 2) != 1 {
		t.Fatal("Wrong with the case : srq.GetValue(0, 2)")
	}
	srq.UpdateSubrectangle(0, 0, 3, 2, 5)
	if srq.GetValue(0, 2) != 5 {
		t.Fatal("Wrong with the case : srq.GetValue(0, 5)")
	}
	if srq.GetValue(3, 1) != 5 {
		t.Fatal("Wrong with the case : srq.GetValue(3, 1)")
	}
	srq.UpdateSubrectangle(3, 0, 3, 2, 10)
	if srq.GetValue(3, 1) != 10 {
		t.Fatal("Wrong with the case : srq.GetValue(3, 1)")
	}
	if srq.GetValue(0, 2) != 5 {
		t.Fatal("Wrong with the case : srq.GetValue(0, 5)")
	}
}

func TestSubrectangleQueries1(t *testing.T) {

	srq := ConstructorSubrectangleQueries([][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}})

	if srq.GetValue(0, 0) != 1 {
		t.Fatal("Wrong with the case : srq.GetValue(0, 0)")
	}
	srq.UpdateSubrectangle(0, 0, 2, 2, 100)
	if srq.GetValue(0, 0) != 100 {
		t.Fatal("Wrong with the case : srq.GetValue(0, 0)")
	}
	if srq.GetValue(2, 2) != 100 {
		t.Fatal("Wrong with the case : srq.GetValue(2, 2)")
	}
	srq.UpdateSubrectangle(1, 1, 2, 2, 20)
	if srq.GetValue(2, 2) != 20 {
		t.Fatal("Wrong with the case : srq.GetValue(2, 2)")
	}
}
