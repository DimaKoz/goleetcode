package easy

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement Queue using Stacks.
Memory Usage: 1.9 MB, less than 78.95% of Go online submissions for Implement Queue using Stacks.
*/

type MyQueue struct {
	stack []int
}

func Constructor() MyQueue {
	q := MyQueue{}
	q.stack = make([]int, 0, 0)
	return q
}

func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	result := this.stack[0]
	this.stack = this.stack[1:len(this.stack)] // Pop
	return result
}

func (this *MyQueue) Peek() int {
	return this.stack[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
