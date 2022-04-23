package easy

import "bytes"

/*
Constraints:

1 <= s.length <= 1000
s consists of digits and the '#' letter.
s will be a valid string such that mapping is always possible.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Decrypt String from Alphabet to Integer Mapping.
Memory Usage: 1.9 MB, less than 90.91% of Go online submissions for Decrypt String from Alphabet to Integer Mapping.
*/
func freqAlphabets(s string) string {
	dict := map[string]string{
		"1":  "a",
		"2":  "b",
		"3":  "c",
		"4":  "d",
		"5":  "e",
		"6":  "f",
		"7":  "g",
		"8":  "h",
		"9":  "i",
		"10": "j",
		"11": "k",
		"12": "l",
		"13": "m",
		"14": "n",
		"15": "o",
		"16": "p",
		"17": "q",
		"18": "r",
		"19": "s",
		"20": "t",
		"21": "u",
		"22": "v",
		"23": "w",
		"24": "x",
		"25": "y",
		"26": "z",
	}
	l := len(s)

	result := bytes.Buffer{}
	v := ""
	i2 := 0
	for i := 0; i < l; i++ {
		i2 = i + 2
		if i2 < l && s[i2] == '#' {
			v, _ = dict[s[i:i2]]
			i = i2
		} else {
			v, _ = dict[string(s[i])]
		}
		result.WriteString(v)
	}
	return result.String()
}
