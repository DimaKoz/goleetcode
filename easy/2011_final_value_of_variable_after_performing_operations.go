package easy

/*
https://leetcode.com/problems/final-value-of-variable-after-performing-operations/

Constraints:

1 <= operations.length <= 100
operations[i] will be either "++X", "X++", "--X", or "X--".
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Final Value of Variable After Performing Operations.
Memory Usage: 3.4 MB, less than 75.91% of Go online submissions for Final Value of Variable After Performing Operations.
*/
func finalValueAfterOperations(operations []string) int {
	value := 0
	operation := ""
	for _, operation = range operations {
		switch operation[1] {
		case '-':
			value--

		default:
			value++
		}
	}
	return value
}
