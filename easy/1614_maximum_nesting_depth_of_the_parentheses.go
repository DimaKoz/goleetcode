package easy

/*
https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/

Constraints:

1 <= s.length <= 100
s consists of digits 0-9 and characters '+', '-', '*', '/', '(', and ')'.
It is guaranteed that parentheses expression s is a VPS.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Maximum Nesting Depth of the Parentheses.
Memory Usage: 1.8 MB, less than 82.61% of Go online submissions for Maximum Nesting Depth of the Parentheses.
*/

/* maxDepth -> maxDepthParentheses renamed*/
func maxDepthParentheses(s string) int {
	depth, open := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			open++
			if open > depth {
				depth = open
			}
		} else if s[i] == ')' {
			open--
		}
	}
	return depth
}
