package easy

import "testing"

func TestDestCity(t *testing.T) {
	type Case struct {
		paths  [][]string
		expect string
	}
	cases := []Case{
		{[][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}, "Sao Paulo"},
		{[][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}, "A"},
	}

	for _, item := range cases {
		if got := destCity(item.paths); got != item.expect {
			t.Fatalf("Wrong with %v", item)
		}
	}
}
