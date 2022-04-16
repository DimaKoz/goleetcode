package easy

import "testing"

func TestLargestPerimeter(t *testing.T) {
	type Case struct {
		income []int
		expect int
	}
	cases := []Case{
		{income: []int{2,1,2}, expect: 5},
		{income: []int{1,2,1}, expect: 0},
	}

	for _, item := range cases {
		if got := largestPerimeter(item.income); got != item.expect {
			t.Fatalf("Wrong with %v", item)
		}
	}
}
