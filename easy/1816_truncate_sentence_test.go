package easy

import "testing"

func TestTruncateSentence(t *testing.T) {
	type Case struct {
		s      string
		k      int
		expect string
	}
	cases := []Case{
		{"Hello how are you Contestant", 4, "Hello how are you"},
		{"What is the solution to this problem", 4, "What is the solution"},
		{"chopper is not a tanuki", 5, "chopper is not a tanuki"},
	}

	for _, item := range cases {
		if got := truncateSentence(item.s, item.k); got != item.expect {
			t.Errorf("Wrong with the case : %v", item)
		}
	}
}
