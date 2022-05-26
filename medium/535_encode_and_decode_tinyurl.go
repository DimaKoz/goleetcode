package medium

/*
https://leetcode.com/problems/encode-and-decode-tinyurl/

Constraints:

1 <= url.length <= 10^4
url is guranteed to be a valid URL.
*/

/*
Runtime: 4 ms, faster than 48.93% of Go online submissions for Encode and Decode TinyURL.
Memory Usage: 6.3 MB, less than 5.26% of Go online submissions for Encode and Decode TinyURL.
*/

import (
	"bytes"
	"math/rand"
	"time"
)

type Codec struct {
	shorts map[string]string
	longs  map[string]string
}

// from Constructor -> ConstructorCodec
func ConstructorCodec() Codec {
	c := Codec{}
	c.shorts = make(map[string]string)
	c.longs = make(map[string]string)
	return c
}

// Encodes a URL to a shortened URL.
// from encode -> encodeCodec
func (this *Codec) encodeCodec(longUrl string) string {
	encodeDictionary := "qwertyuiopasdfghjklzxcvbnm1234567890"
	b := bytes.Buffer{}
	for true {
		n := getRandomNumberForDictionary(len(encodeDictionary) - 1)
		b.WriteByte(encodeDictionary[n])
		if _, found := this.shorts[b.String()]; !found {
			break
		}
	}
	this.shorts[b.String()] = longUrl
	this.longs[longUrl] = b.String()
	return "http://tinyurl.com/" + b.String()
}

var randomcounter int64

func getRandomNumberForDictionary(max int) int {
	randomcounter++
	s1 := rand.NewSource(time.Now().UnixNano() + randomcounter)
	r1 := rand.New(s1)
	return r1.Intn(max)
}

// Decodes a shortened URL to its original URL.
// from decode -> decodeCodec
func (this *Codec) decodeCodec(shortUrl string) string {
	search := shortUrl[len("http://tinyurl.com/"):len(shortUrl)]
	if result, found := this.shorts[search]; found {
		return result
	} else {
		return ""
	}
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
