package stack

import "testing"

func TestStackLen(t *testing.T) {
	var myStack Stack
	myStack.Push(1)
	myStack.Push("test")
	if myStack.Len() == 2 {
		t.Log("Pass Stack.Len")
	} else {
		t.Error("Failed Stack.Len")
	}
}

func TestStackIsEmpty(t *testing.T) {
	var mStack Stack
	if mStack.IsEmpty() {
		t.Log("Pass Stack.IsEmpty")
	} else {
		t.Error("Failed Stack.IsEmpty")
	}
}

func TestStackCap(t *testing.T) {
	myStack := make(Stack, 3)
	if myStack.Cap() == 3 {
		t.Log("Pass Stack.Cap")
	} else {
		t.Error("Failed Stack.Cap")
	}
}

func TestStackPush(t *testing.T) {
	var mStack Stack
	mStack.Push(3)
	if mStack.Len() == 1 {
		t.Log("Pass Stack.Push")
	} else {
		t.Error("Failed Stack.Push")
	}
}

func TestStackTop(t *testing.T) {
	var mStack Stack
	if _, err := mStack.Top(); err == nil {
		t.Error("Failed Stack.Top")
	}
	mStack.Push(3)
	if value, _ := mStack.Top(); value == 3 {
		t.Log("Pass Stack.Top")
	} else {
		t.Errorf("Failed Stack.Top, value is %d", value)
	}
}

func TestStackPop(t *testing.T) {
	var mStack Stack
	if _, err := mStack.Pop(); err == nil {
		t.Error("Failed Stack.Top")
	}
	mStack.Push("test")
	mStack.Push(3)
	if value, _ := mStack.Pop(); value == 3 && mStack.Len() == 1 {
		t.Log("Pass Stack.Pop")
	} else {
		t.Errorf("Failed Stack.Pop, value is %d, len is %d", value, mStack.Len())
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
