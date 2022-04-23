package easy

import "testing"

func TestReverseString(t *testing.T) {
	type Case struct {
		s      []byte
		expect []byte
	}
	cases := []Case{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
	}

	for _, item := range cases {
		reverseString(item.s)
		for i := 0; i < len(item.s); i++ {
			if got := item.s[i]; got != item.expect[i] {
				t.Errorf("reverseString Wrong with the case : %v, got:%s, expect%s", item, string(got), string(item.expect[i]))
			}
		}
	}

}
