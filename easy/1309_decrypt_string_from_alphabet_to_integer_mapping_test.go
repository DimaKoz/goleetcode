package easy

import "testing"

func TestFreqAlphabets(t *testing.T) {
	type Case struct {
		s      string
		expect string
	}
	cases := []Case{
		{"10#11#12", "jkab"},
		{"1326#", "acz"},
	}

	for _, item := range cases {
		if got := freqAlphabets(item.s); got != item.expect {
			t.Errorf("Wrong with the case : %v, got:%s", item, got)
		}
	}
}
