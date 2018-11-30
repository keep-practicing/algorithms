// 单链表实现

package singlelinkedlist

// Node is an element of a linked list.
type Node struct {
	Value interface{}
	next  *Node
}

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type List struct {
	Head *Node
	len  int
}

// Init initializes or clears list l.
func (l *List) Init() *List {
	l.Head = &Node{}
	l.len = 0
	return l
}

// New returns an initialized list.
func New() *List {
	return new(List).Init()
}

// IsEmpty judge the list is empty
func (l *List) IsEmpty() bool {
	return l.Head.next == nil
}

// Length size of list
func (l *List) Length() int {
	return l.len
}

// Clear clear list
func (l *List) Clear() {
	l.Init()
}

// Insert inserts n, increments l.len, and returns e.
func (l *List) Insert(i int, n *Node) bool {
	if i < 1 || i > l.len+1 {
		return false
	}

	current := l.Head
	for k := 1; k < i; k++ {
		current = current.next
	}

	n.next = current.next
	current.next = n
	l.len++
	return false
}

// Delete delete i from list.
func (l *List) Delete(i int) (r interface{}, ok bool) {
	if i < 1 || i > l.len {
		return
	}

	current := l.Head
	var before *Node

	for k := 1; k <= i; k++ {
		before, current = current, current.next
	}

	before.next = current.next
	r = current.Value
	current = nil
	l.len--
	return r, true
}

// GetNode get node from list by index.
func (l *List) GetNode(index int) (r Node, ok bool) {
	if index < 1 || index > l.len {
		return
	}

	current := l.Head
	for k := 1; k <= index; k++ {
		current = current.next
	}

	r = Node{Value: current.Value}
	return r, true
}

// GetIndex get index if Value is x.
func (l *List) GetIndex(x interface{}) (index int) {
	for current := l.Head; current.next != nil; {
		current = current.next
		if current.Value == x {
			break
		}
		index++
	}
	return
}
