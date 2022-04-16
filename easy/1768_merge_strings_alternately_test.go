package easy

import "testing"

func TestMergeAlternately(t *testing.T) {
	type Case struct {
		s1     string
		s2     string
		expect string
	}
	cases := []Case{
		{"abcd","pq", "apbqcd"},
		{"ab", "pqrs", "apbqrs"},
		{"abc", "pqr", "apbqcr"},
	}

	for _, item := range cases {
		if got := mergeAlternately(item.s1, item.s2); got != item.expect {
			t.Errorf("Wrong with the case : %v, got:%s", item, got)
		}
	}
}
