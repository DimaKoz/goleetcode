package easy

import "bytes"

/*
https://leetcode.com/problems/remove-outermost-parentheses/

Constraints:

1 <= s.length <= 10^5
s[i] is either '(' or ')'.
s is a valid parentheses string.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Outermost Parentheses.
Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Remove Outermost Parentheses.
*/

func removeOuterParentheses(s string) string {
	const NoInd = -1
	outerIndStart := NoInd
	openCounter, closeCounter := 0, 0
	b := bytes.Buffer{}
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if openCounter == 0 {
				outerIndStart = i
			}
			openCounter++
		}
		if s[i] == ')' {
			closeCounter++
			if openCounter == closeCounter {
				closeCounter, openCounter = 0, 0
				b.WriteString(s[outerIndStart+1 : i])
				outerIndStart = NoInd
			}
		}
	}
	return b.String()
}
