package queue

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue()

	if q.Poll() != nil {
		t.Error("test failed")
	}

	q.Offer(2)
	q.Offer("hello, world")

	if q.Size() != 2 {
		t.Error("test failed")
	}

	if q.Poll() != 2 {
		t.Error("test failed")
	}

	q.Clear()

	if q.Clear() != false {
		t.Error("test failed")
	}

	if q.IsEmpty() != true {
		t.Error("test  failed")
	}
}
