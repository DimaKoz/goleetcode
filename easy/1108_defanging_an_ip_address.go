package easy

/*
https://leetcode.com/problems/defanging-an-ip-address/

Constraints:

The given address is a valid IPv4 address.
*/

/*
Runtime: 2 ms, faster than 32.54% of Go online submissions for Defanging an IP Address.
Memory Usage: 1.9 MB, less than 26.19% of Go online submissions for Defanging an IP Address.
*/

import "bytes"

func defangIPaddr(address string) string {
	b := bytes.Buffer{}
	b.Grow(21 /*max output ipv4 len*/)
	lenA := len(address)
	var sym byte
	for i := 0; i < lenA; i++ {
		sym = address[i]
		if sym == '.' {
			b.WriteString("[.]")
		} else {
			b.WriteByte(sym)
		}
	}
	return b.String()
}
