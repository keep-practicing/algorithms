package addTwoNumbers

import "testing"

func TestAddTwoNumbers1(t *testing.T) {
	var (
		l1Vals   = [3]int{2, 4, 3}
		l2Vals   = [3]int{5, 6, 4}
		l1       = ListNode{Val: -1}
		l2       = l1
		cur1     = &l1
		cur2     = &l2
		addLi1   [3]int
		testFunc = [2]func(l1 *ListNode, l2 *ListNode) *ListNode{addTwoNumbers1, addTwoNumbers2}
	)

	expectedLi := [3]int{7, 0, 8}

	for i := 0; i < len(l1Vals); i++ {
		cur1.Next = &ListNode{Val: l1Vals[i]}
		cur1 = cur1.Next
	}
	for i := 0; i < len(l2Vals); i++ {
		cur2.Next = &ListNode{Val: l2Vals[i]}
		cur2 = cur2.Next
	}

	for i := 0; i < 2; i++ {
		addHead := testFunc[i](l1.Next, l2.Next)

		for i := 0; addHead != nil; i++ {
			addLi1[i] = addHead.Val
			addHead = addHead.Next
		}

		if addLi1 != expectedLi {
			t.Error("测试不通过")
		}
	}
}
