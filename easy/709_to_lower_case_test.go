package easy

import "testing"

func TestToLowerCase(t *testing.T) {
	type Case struct {
		input  string
		expect string
	}
	cases := []Case{
		{input: "Hello", expect: "hello"},
		{input: "here", expect: "here"},
		{input: "LOVELY", expect: "lovely"},
	}
	for index, item := range cases {
		if got := toLowerCase(item.input); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
