package easy

/*
https://leetcode.com/problems/count-items-matching-a-rule/

Constraints:

1 <= items.length <= 10^4
1 <= typei.length, colori.length, namei.length, ruleValue.length <= 10
ruleKey is equal to either "type", "color", or "name".
All strings consist only of lowercase letters.
*/

/*
Runtime: 29 ms, faster than 91.09% of Go online submissions for Count Items Matching a Rule.
Memory Usage: 7 MB, less than 80.20% of Go online submissions for Count Items Matching a Rule.
*/

func countMatches(items [][]string, ruleKey string, ruleValue string) int {

	lenS := len(items)
	var ind, counter, i int
	if ruleKey == "color" {
		ind = 1
	} else if ruleKey == "name" {
		ind = 2
	}

	for i = 0; i < lenS; i++ {
		if items[i][ind] == ruleValue {
			counter++
		}
	}
	return counter
}
