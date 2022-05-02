package easy

import "testing"

func TestFinalValueAfterOperations(t *testing.T) {
	type Case struct {
		operations []string
		expect     int
	}
	cases := []Case{
		{[]string{"--X", "X++", "X++"}, 1},
		{[]string{"++X", "++X", "X++"}, 3},
		{[]string{"X++", "++X", "--X", "X--"}, 0},
	}
	for index, item := range cases {
		if got := finalValueAfterOperations(item.operations); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
