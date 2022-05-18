package easy

import "testing"

func TestUniqueMorseRepresentations(t *testing.T) {
	type Case struct {
		words  []string
		expect int
	}
	cases := []Case{
		{[]string{"gin", "zen", "gig", "msg"}, 2},
		{[]string{"a"}, 1},
	}

	for index, item := range cases {
		if got := uniqueMorseRepresentations(item.words); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
