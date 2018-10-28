package singleLinkedList

import "testing"

func TestReverseSingleLinkedList(t *testing.T) {
	head := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	listVal := [...]int{5, 4, 3, 2, 1}

	reverseHead := reverseSingleLinkedList(&head)

	for _, j := range listVal {
		if j != reverseHead.Val {
			t.Error("测试不通过")
		}
		reverseHead = reverseHead.Next
	}

	head1 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	reverseHead = reverseSingleLinkedList(&head1)
	for _, j := range listVal {
		if j != reverseHead.Val {
			t.Error("测试不通过")
		}
		reverseHead = reverseHead.Next
	}
}
