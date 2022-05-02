package easy

import "testing"

func TestMostWordsFound(t *testing.T) {
	type Case struct {
		sentences []string
		expect    int
	}
	cases := []Case{
		{[]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}, 6},
		{[]string{"please wait", "continue to fight", "continue to win"}, 3},
	}
	for index, item := range cases {
		if got := mostWordsFound(item.sentences); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
