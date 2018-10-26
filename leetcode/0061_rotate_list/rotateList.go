package rotateList

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var (
		l2 = ListNode{Val: 2, Next: nil}
		l1 = ListNode{Val: 1, Next: &l2}
		l0 = ListNode{Val: 0, Next: &l1}
	)
	rotateRight(&l0, 4)
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	var (
		listSize   int
		count      int
		p          = head
		rotateNode = head
	)

	for p != nil && count < k {
		p = p.Next
		listSize++
		count++
	}

	if p == nil {
		k = k % listSize
		if k == 0 {
			return head
		}
		p = head
		for count = 0; count < k; count++ {
			p = p.Next
		}
	}

	for p.Next != nil {
		rotateNode = rotateNode.Next
		p = p.Next
	}

	p.Next = head
	head = rotateNode.Next
	rotateNode.Next = nil
	return head
}
