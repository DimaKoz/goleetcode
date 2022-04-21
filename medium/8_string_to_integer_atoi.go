package medium

/*

https://leetcode.com/problems/string-to-integer-atoi/

Constraints:

0 <= s.length <= 200
s consists of English letters (lower-case and upper-case), digits (0-9), ' ', '+', '-', and '.'.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for String to Integer (atoi).
Memory Usage: 2.5 MB, less than 7.95% of Go online submissions for String to Integer (atoi).
*/
func myAtoi(s string) int {
	isNegative := false
	lenS := len(s)
	number := make([]int, 0, lenS)
	var c uint8
	var prev uint8
	var haveDigits = false
	for i := 0; i < lenS; i++ {
		c = s[i]
		if c == '-' || c == '+' {
			if haveDigits {
				break
			}
			haveDigits = true
			prev = c
			continue
		}
		if c == ' ' {
			if haveDigits {
				break
			}
			continue
		}
		if '0' <= c && c <= '9' {
			haveDigits = true
			if prev == '-' {
				isNegative = true
				prev = 0
			}
			if len(number) == 0 && c == '0' {
				continue
			}
			number = append(number, int(c-'0'))
			continue
		}
		break
	}
	if len(number) == 0 {
		return 0
	}

	// clamp
	if len(number) > 10 {
		if isNegative {
			return -2147483648 // min int32
		} else {
			return 2147483647 // max int32
		}
	}

	result := sliceToIntMyAtoi(number)

	if isNegative {
		if result > 2147483648 { // clamp
			return -2147483648 // min int32
		}
		return -int(result)
	} else {
		if result > 2147483647 { // clamp
			return 2147483647 // max int32
		}
		return int(result)
	}
}

func sliceToIntMyAtoi(s []int) int64 {
	res := int64(0)
	op := int64(1)
	for i := len(s) - 1; i >= 0; i-- {
		res += int64(s[i]) * op
		op *= 10
	}
	return res
}
