package singlelinkedlist

import "testing"

func TestReverseSingleLinkedList(t *testing.T) {
	listVal := [...]int{5, 4, 3, 2, 1}

	// 方法二测试
	head := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	reverseHead := reverseSingleLinkedList(&head)
	for _, j := range listVal {
		if j != reverseHead.Val {
			t.Error("测试不通过")
		}
		reverseHead = reverseHead.Next
	}

	// 方法三测试
	head1 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	reverseHead = reverseSingleLinkedList2(&head1)
	for _, j := range listVal {
		if j != reverseHead.Val {
			t.Error("测试不通过")
		}
		reverseHead = reverseHead.Next
	}

	// 方法四测试
	head2 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	reverseHead = reverseSingleLinkedList3(&head2)
	for _, j := range listVal {
		if j != reverseHead.Val {
			t.Error("测试不通过")
		}
		reverseHead = reverseHead.Next
	}
}
