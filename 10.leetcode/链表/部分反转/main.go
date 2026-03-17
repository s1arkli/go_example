package main

/*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	curr := head
	var (
		l, r   *ListNode
		lr, rl *ListNode
	)
	for i := 1; i <= right; i++ {
		if i == left-1 {
			l = curr.Next
			lr = curr
			curr.Next = nil
		}
		if i == right {
			r = curr
			rl = curr.Next
			r.Next = nil
		}
		curr = curr.Next
	}
	reverse(l)
	lr.Next = r
	l.Next = rl
	return head
}

func reverse(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
