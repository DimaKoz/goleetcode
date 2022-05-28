package easy

/*
https://leetcode.com/problems/design-an-ordered-stream/

Constraints:

1 <= n <= 1000
1 <= id <= n
value.length == 5
value consists only of lowercase letters.
Each call to insert will have a unique id.
Exactly n calls will be made to insert.
*/

/*
Runtime: 69 ms, faster than 96.19% of Go online submissions for Design an Ordered Stream.
Memory Usage: 7.2 MB, less than 88.57% of Go online submissions for Design an Ordered Stream.
*/

type OrderedStream struct {
	last   int
	values []string
}

//Constructor renamed -> ConstructorOrderedStream
func ConstructorOrderedStream(n int) OrderedStream {
	stream := OrderedStream{}
	stream.values = make([]string, n+1)
	return stream
}

//Insert renamed -> InsertOrderedStream
func (this *OrderedStream) InsertOrderedStream(idKey int, value string) []string {
	this.values[idKey-1] = value

	result := make([]string, 0)
	for i := this.last; i < len(this.values); i++ {
		if this.values[i] == "" {
			break
		}
		result = append(result, this.values[i])
		this.last = i + 1
	}
	return result
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
