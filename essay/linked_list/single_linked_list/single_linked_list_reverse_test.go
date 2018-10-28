package singleLinkedList

import "testing"

func TestReverseSingleLinkedList(t *testing.T) {
	head := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	reverseHead := reverseSingleLinkedList(&head)
	listVal := [...]int{5, 4, 3, 2, 1}
	for _, j := range listVal {
		if j != reverseHead.Val {
			t.Error("测试不通过")
		}
		reverseHead = reverseHead.Next
	}
}
