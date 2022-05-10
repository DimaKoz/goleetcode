package easy

import "testing"

func TestCountMatches(t *testing.T) {
	type Case struct {
		items     [][]string
		ruleKey   string
		ruleValue string
		expect    int
	}

	cases := []Case{
		{
			items: [][]string{
				{"phone", "blue", "pixel"},
				{"computer", "silver", "lenovo"},
				{"phone", "gold", "iphone"},
			},
			ruleKey:   "color",
			ruleValue: "silver",
			expect:    1},

		{
			items: [][]string{
				{"phone", "blue", "pixel"},
				{"computer", "silver", "lenovo"},
				{"phone", "gold", "iphone"},
			},
			ruleKey:   "type",
			ruleValue: "phone",
			expect:    2},
	}

	for _, item := range cases {
		if got := countMatches(item.items, item.ruleKey, item.ruleValue); got != item.expect {
			t.Errorf("Wrong with the case : %v", item)
		}
	}
}
