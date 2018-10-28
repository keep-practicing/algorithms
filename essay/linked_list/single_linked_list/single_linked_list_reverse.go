package singleLinkedList

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
