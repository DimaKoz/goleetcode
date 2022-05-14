package easy

import "testing"

func TestMinMovesToSeat(t *testing.T) {
	type Case struct {
		seats    []int
		students []int
		expect   int
	}
	cases := []Case{
		{[]int{3, 20, 17, 2, 12, 15, 17, 4, 15, 20}, []int{10, 13, 14, 15, 5, 2, 3, 14, 3, 18}, 28},
		{[]int{3, 1, 5}, []int{2, 7, 4}, 4},
		{[]int{4, 1, 5, 9}, []int{1, 3, 2, 6}, 7},
		{[]int{2, 2, 6, 6}, []int{1, 3, 2, 6}, 4},
	}
	for index, item := range cases {
		if got := minMovesToSeat(item.seats, item.students); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
