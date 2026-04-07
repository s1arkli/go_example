package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	a := 0

	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + a
		if sum >= 10 {
			a = 1
		} else {
			a = 0
		}

		cur.Next = &ListNode{
			Val: sum % 10,
		}
		cur = cur.Next
	}
	if a > 0 {
		cur.Next = &ListNode{
			Val: a,
		}
	}

	return dummy.Next
}
