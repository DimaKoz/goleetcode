package easy

import "testing"

func TestInterpret(t *testing.T) {
	type Case struct {
		command string
		expect  string
	}
	cases := []Case{
		{command: "GG()", expect: "GGo"},
		{command: "G()(al)", expect: "Goal"},
		{command: "G()()()()(al)", expect: "Gooooal"},
	}

	for _, item := range cases {
		if got := interpret(item.command); got != item.expect {
			t.Errorf("Wrong with the case : %v, got:%s", item, got)
		}
	}
}
