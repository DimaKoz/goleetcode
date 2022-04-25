package easy

import "testing"

func TestMajorityElement(t *testing.T) {
	type Case struct {
		income []int
		expect int
	}
	cases := []Case{
		{income: []int{2, 2, 1, 1, 1, 2, 2}, expect: 2},
		{income: []int{3, 2, 3}, expect: 3},
	}

	for _, item := range cases {
		if got := majorityElement(item.income); got != item.expect {
			t.Fatalf("Wrong with %v", item)
		}
	}
}
