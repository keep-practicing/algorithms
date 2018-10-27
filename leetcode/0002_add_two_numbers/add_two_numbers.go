package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法一，暴力解法
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{Val: -1}
	cur := &head
	carry := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		val := sum % 10
		carry = sum / 10
		node := ListNode{Val: val}
		cur.Next = &node
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if carry == 0 {
		if l1 != nil {
			cur.Next = l1
		}
		if l2 != nil {
			cur.Next = l2
		}
		return head.Next
	} else {
		for l1 != nil {
			sum := l1.Val + carry
			val := sum % 10
			carry = sum / 10
			node := ListNode{Val: val}
			cur.Next = &node
			cur = cur.Next
			l1 = l1.Next
		}
		for l2 != nil {
			sum := l2.Val + carry
			val := sum % 10
			carry = sum / 10
			node := ListNode{Val: val}
			cur.Next = &node
			cur = cur.Next
			l2 = l2.Next
		}
		for carry != 0 {
			val := carry % 10
			carry = carry / 10
			node := ListNode{Val: val}
			cur.Next = &node
			cur = cur.Next
		}
	}
	return head.Next
}

// 解法二，递归
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil && l2 != nil {
		return l2
	} else if l1 != nil && l2 == nil {
		return l1
	}

	sum := l1.Val + l2.Val
	next := addTwoNumbers2(l1.Next, l2.Next)

	if sum >= 10 {
		carry := sum / 10
		sum %= 10
		next = addTwoNumbers2(next, &ListNode{Val: carry})
	}
	return &ListNode{Val: sum, Next: next}
}
