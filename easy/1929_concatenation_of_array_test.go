package easy

import "testing"

func TestGetConcatenation(t *testing.T) {
	type Case struct {
		input  []int
		expect []int
	}
	cases := []Case{
		{input: []int{1, 3, 2, 1}, expect: []int{1, 3, 2, 1, 1, 3, 2, 1}},
	}
	for index, item := range cases {

		if got := getConcatenation(item.input); len(got) != len(item.expect) {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		} else {
			for i := 0; i < len(got); i++ {
				if got[i] != item.expect[i] {
					t.Fatalf("Wrong with the case : %v, index: %d at:%d", item, index, i)
				}
			}
		}
	}
}
