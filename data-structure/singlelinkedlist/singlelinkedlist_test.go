package singlelinkedlist

import (
	"testing"
)

func TestNewList(t *testing.T) {
	list := New()

	if list.Head == nil {
		t.Error("test failed")
	}

	if list.Head.next != nil {
		t.Error("test failed")
	}
}

func TestListInsert(t *testing.T) {
	list := New()
	list.Insert(1, &Node{Value: 23})
	if list.Length() != 1 {
		t.Error("test failed")
	}

	if ok := list.Insert(-1, &Node{Value: 23}); ok != false {
		t.Error("test failed")
	}

	if ok := list.Insert(10, &Node{Value: 23}); ok != false {
		t.Error("test failed")
	}
}

func TestListGetNode(t *testing.T) {
	list := New()

	list.Insert(1, &Node{Value: 23})
	list.Insert(1, &Node{Value: "hello, world"})
	list.Insert(1, &Node{Value: 23.2})
	list.Insert(1, &Node{Value: 100})

	if _, ok := list.GetNode(0); ok {
		t.Error("test failed")
	}
	if _, ok := list.GetNode(5); ok {
		t.Error("test failed")
	}

	if n, ok := list.GetNode(2); !ok || n.Value != "hello, world" {
		t.Error("test failed")
	}
}

func TestListDelete(t *testing.T) {
	list := New()
	list.Insert(1, &Node{Value: 23})
	list.Insert(1, &Node{Value: "hello, world"})
	list.Insert(1, &Node{Value: 23.2})
	list.Insert(1, &Node{Value: 100})

	if n, ok := list.Delete(2); !ok || n != "hello, world" {
		t.Error("test failed")
	}

	if _, ok := list.Delete(0); ok {
		t.Error("test failed")
	}

	if _, ok := list.Delete(5); ok {
		t.Error("test failed")
	}
}

func TestListGetIndex(t *testing.T) {
	list := New()
	list.Insert(1, &Node{Value: 23})
	list.Insert(1, &Node{Value: "hello, world"})
	list.Insert(1, &Node{Value: 23.2})
	list.Insert(1, &Node{Value: 100})

	if index := list.GetIndex("hello, world"); index != 2 {
		t.Error("test failed")
	}
}

func TestListClear(t *testing.T) {
	list := New()
	list.Insert(1, &Node{Value: 23})
	list.Insert(1, &Node{Value: "hello, world"})
	list.Insert(1, &Node{Value: 23.2})
	list.Insert(1, &Node{Value: 100})

	list.Clear()

	if list.len != 0 || list.Head.next != nil {
		t.Error("test failed")
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
