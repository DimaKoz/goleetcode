package easy

/*
Constraints:

m == accounts.length
n == accounts[i].length
1 <= m, n <= 50
1 <= accounts[i][j] <= 100
*/
func maximumWealth(accounts [][]int) int {

	var customer []int
	var max, sum, acc int
	for _, customer = range accounts {
		sum = 0
		for _, acc = range customer {
			sum += acc
		}
		if max < sum {
			max = sum
		}
	}
	return max
}
