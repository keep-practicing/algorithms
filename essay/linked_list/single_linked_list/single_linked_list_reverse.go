package singlelinkedlist

// ListNode is node of linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法二
func reverseSingleLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		q = head.Next
		r *ListNode
	)

	head.Next = nil
	for q != nil {
		r = q.Next
		q.Next = head
		head = q
		q = r
	}
	return head
}

// 方法三
func reverseSingleLinkedList2(head *ListNode) *ListNode {
	var (
		p = head.Next
		q *ListNode
	)

	for p.Next != nil {
		q = p.Next
		p.Next = q.Next
		q.Next = head.Next
		head.Next = q
	}

	p.Next = head      // 单链表成环
	head = p.Next.Next // 新head变为原head的next  head=q
	p.Next.Next = nil  // 断链表
	return head
}

// 方法四 递归
func reverseSingleLinkedList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head //链表为空直接返回，head.Next 为空时递归基
	}

	newHead := reverseSingleLinkedList3(head.Next) // 递归调用
	head.Next.Next = head                          // 分析两个节点的情况，形成环
	head.Next = nil                                // 断开链表
	return newHead
}
