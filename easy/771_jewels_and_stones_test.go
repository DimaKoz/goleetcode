package easy

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	type Case struct {
		j      string
		s      string
		expect int
	}
	cases := []Case{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}

	for index, item := range cases {
		if got := numJewelsInStones(item.j, item.s); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
