package queue

import (
	"fmt"
)

// Element iof Queue
type Element interface{}

// Queue data structure
type Queue interface {
	Offer(e Element) // 向队列中添加元素
	Poll() Element   //移除队列中最前面的元素
	Clear() bool     // clear queue
	Size() int       // 获取队列的元素个数
	IsEmpty() bool   // 判断队列是否为空
}

// SliceEntry is queue head
type SliceEntry struct {
	element []Element
}

// NewQueue func is to init queue
func NewQueue() *SliceEntry {
	return &SliceEntry{}
}

// Offer 向队列中添加元素
func (entry *SliceEntry) Offer(e Element) {
	entry.element = append(entry.element, e)
}

// IsEmpty   队列是否为空
func (entry *SliceEntry) IsEmpty() bool {
	if len(entry.element) == 0 {
		return true
	}
	return false
}

// Poll 移除队列中最前面的元素
func (entry *SliceEntry) Poll() Element {
	if entry.IsEmpty() {
		fmt.Println("queue is empty")
		return nil
	}
	firstElement := entry.element[0]
	entry.element = entry.element[1:]
	return firstElement
}

// Clear 清空队列
func (entry *SliceEntry) Clear() bool {
	if entry.IsEmpty() {
		fmt.Println("queue is empty")
		return false
	}
	for i := 0; i < entry.Size(); i++ {
		entry.element[i] = nil
	}
	entry.element = nil
	return true
}

// Size 返回元素个数
func (entry *SliceEntry) Size() int {
	return len(entry.element)
}
