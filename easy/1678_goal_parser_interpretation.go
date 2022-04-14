package easy

import (
	"bytes"
	"strings"
)

/*
Constraints:

1 <= command.length <= 100
command consists of "G", "()", and/or "(al)" in some order.

*/

func interpret(command string) string {
	return interpret2Variant(command)
}

/*
Runtime: 2 ms, faster than 36.80% of Go online submissions for Goal Parser Interpretation.
Memory Usage: 1.9 MB, less than 98.40% of Go online submissions for Goal Parser Interpretation.
*/
func interpret1Variant(command string) string {
	command = strings.Replace(command, "()", "o", -1)
	command = strings.Replace(command, "(al)", "al", -1)
	return command
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Goal Parser Interpretation.
Memory Usage: 1.9 MB, less than 98.40% of Go online submissions for Goal Parser Interpretation.
*/
func interpret2Variant(command string) string {
	var b bytes.Buffer
	l := len(command)
	b.Grow(l)
	for i := 0; i < l; i++ {
		if command[i] == '(' {
			if i+1 < l && command[i+1] == ')' {
				b.WriteByte('o')
				i++
			}
			continue
		} else if command[i] == ')' {
			continue
		}
		b.WriteByte(command[i])
	}
	return b.String()
}
