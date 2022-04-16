package easy

import "testing"

func TestMyQueue(t *testing.T) {
	myQueue := Constructor()
	myQueue.Push(1)
	myQueue.Push(2)
	gotPeek := myQueue.Peek()
	if gotPeek != 1 { //FIFO
		t.Errorf("Wrong with myQueue.Peek(), got:%d", gotPeek)
	}
	gotPop := myQueue.Pop()
	if gotPop != 1 { //FIFO
		t.Errorf("Wrong with myQueue.Pop(), got:%d", gotPop)
	}

	if myQueue.Empty() {
		t.Errorf("Wrong with myQueue.Empty()")
	}

	myQueue.Pop()
	if !myQueue.Empty() {
		t.Errorf("Wrong with myQueue.Empty()")
	}
}
