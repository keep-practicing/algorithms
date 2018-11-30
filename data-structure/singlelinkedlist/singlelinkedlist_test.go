package singlelinkedlist

import (
	"os"
	"testing"
)

func TestNewList(t *testing.T) {
	list := New()

	if list.Head == nil {
		t.Error("Failed List.New")
	}

	if list.Head.next != nil {
		t.Error("Failed List.New")
	}
}

func TestListInsert(t *testing.T) {
	list := New()
	list.Insert(1, &Node{Value: 23})
	list.Insert(2, &Node{Value: "hello, world"})
	list.Insert(3, &Node{Value: 23.2})
	list.Insert(4, &Node{Value: 100})

	if list.Length() != 4 {
		t.Error("Failed List.Insert")
	}

	if ok := list.Insert(-1, &Node{Value: 23}); ok != false {
		t.Error("test failed")
	}

	if ok := list.Insert(4, &Node{Value: "你好"}); ok && !(list.GetIndex("你好") == 4 && list.GetIndex(100) == 5) {
		t.Error("Failed List.Insert")
	}

	if ok := list.Insert(10, &Node{Value: 23}); ok != false {
		t.Error("test failed")
	}
}

func TestListGetNode(t *testing.T) {
	list := New()

	list.Insert(1, &Node{Value: 23})
	list.Insert(2, &Node{Value: "hello, world"})
	list.Insert(3, &Node{Value: 23.2})
	list.Insert(4, &Node{Value: 100})

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
	list.Insert(2, &Node{Value: "hello, world"})
	list.Insert(3, &Node{Value: 23.2})
	list.Insert(4, &Node{Value: 100})

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
	list.Insert(2, &Node{Value: "hello, world"})
	list.Insert(3, &Node{Value: 23.2})
	list.Insert(4, &Node{Value: 100})

	if index := list.GetIndex("hello, world"); index != 2 {
		t.Error("test failed")
	}
}

func TestListClear(t *testing.T) {
	list := New()
	list.Insert(1, &Node{Value: 23})
	list.Insert(2, &Node{Value: "hello, world"})
	list.Insert(3, &Node{Value: 23.2})
	list.Insert(4, &Node{Value: 100})

	list.Clear()

	if !list.IsEmpty() || list.Head == nil {
		t.Error("test failed")
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
