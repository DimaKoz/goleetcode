package easy

import "testing"

func TestContainsDuplicate(t *testing.T) {
	type Case struct {
		s      []int
		expect bool
	}
	cases := []Case{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, item := range cases {
		if got := containsDuplicate(item.s); got != item.expect {
			t.Errorf("Wrong with the case : %v, got:%t", item, got)
		}
		if got := containsDuplicate2(item.s); got != item.expect {
			t.Errorf("Wrong with the case : %v, got:%t", item, got)
		}
	}
}
