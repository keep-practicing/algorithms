package singleLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

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
