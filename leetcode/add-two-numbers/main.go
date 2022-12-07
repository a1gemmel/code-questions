package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) Equals(other *ListNode) bool {
	if ln == nil || other == nil {
		return ln == nil && other == nil
	}
	return ln.Val == other.Val && ln.Next.Equals(other.Next)
}

func addTwoNumbersWithCarry(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil {
		if carry == 0 {
			return l2
		}
		l1 = &ListNode{Val: 0}
	}
	if l2 == nil {
		if carry == 0 {
			return l1
		}
		l2 = &ListNode{Val: 0}
	}

	digit := l1.Val + l2.Val + carry

	return &ListNode{Val: digit % 10, Next: addTwoNumbersWithCarry(l1.Next, l2.Next, digit/10)}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersWithCarry(l1, l2, 0)
}
