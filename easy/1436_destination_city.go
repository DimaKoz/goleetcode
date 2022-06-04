package easy

/*
https://leetcode.com/problems/destination-city/

Constraints:

1 <= paths.length <= 100
paths[i].length == 2
1 <= cityAi.length, cityBi.length <= 10
cityAi != cityBi
All strings consist of lowercase and uppercase English letters and the space character.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Destination City.
Memory Usage: 4 MB, less than 75.81% of Go online submissions for Destination City.
*/

func destCity(paths [][]string) string {
	cities := make(map[string]int)
	counter := 0
	for i := 0; i < len(paths); i++ {
		for ii := 0; ii < len(paths[i]); ii++ {
			counter, _ = cities[paths[i][ii]]
			if ii == 0 {
				counter++
			}
			cities[paths[i][ii]] = counter
		}
	}
	for k, v := range cities {
		if v == 0 {
			return k
		}
	}
	return ""
}
