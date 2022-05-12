package easy

import "testing"

func TestXorOperation(t *testing.T) {
	type Case struct {
		n      int
		start  int
		expect int
	}
	cases := []Case{
		{5, 0, 8},
		{4, 3, 8},
	}
	for index, item := range cases {
		if got := xorOperation(item.n, item.start); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d, got:%d", item, index, got)
		}
	}
}
