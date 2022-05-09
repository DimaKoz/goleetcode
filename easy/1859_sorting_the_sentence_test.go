package easy

import "testing"

func TestSortSentence(t *testing.T) {
	type Case struct {
		s      string
		expect string
	}
	cases := []Case{
		{"is2 sentence4 This1 a3", "This is a sentence"},
		{"Myself2 Me1 I4 and3", "Me Myself and I"},
	}
	for index, item := range cases {
		if got := sortSentence(item.s); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
