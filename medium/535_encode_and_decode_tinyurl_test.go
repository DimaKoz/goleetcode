package medium

import "testing"

func TestEncodeCodec(t *testing.T) {
	longUrl := "https://leetcode.com/problems/design-tinyurl"
	obj := ConstructorCodec()
	url := obj.encodeCodec(longUrl)
	got := obj.decodeCodec(url)
	if got != longUrl {
		t.Errorf("Wrong with the case : %v, got:%v", longUrl, got)
	}
}
