package easy

import "testing"

func TestDefangIPaddr(t *testing.T) {
	type Case struct {
		input  string
		expect string
	}
	cases := []Case{
		{input: "1.1.1.1", expect: "1[.]1[.]1[.]1"},
		{input: "255.100.50.0", expect: "255[.]100[.]50[.]0"},
	}

	for index, item := range cases {
		if got := defangIPaddr(item.input); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
