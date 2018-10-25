package merge_two_sorted_lists

import "testing"

func TestMergeTwoLists(t *testing.T) {
	l2 := ListNode{Val: 2, Next: nil}
	head1 := ListNode{Val: 1, Next: &l2}
	head2 := ListNode{Val: 3, Next: nil}

	listNodeArray := [3]*ListNode{&head1, &l2, &head2}

	head := mergeTwoLists(&head1, &head2)

	for _, j := range listNodeArray {
		if j != head {
			t.Error("测试不通过.")
		}
		head = head.Next
	}
}
