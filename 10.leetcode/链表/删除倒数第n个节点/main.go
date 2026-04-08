package main

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd1 删除第n个节点，只需要把第n个节点的前一个节点a.next=a.next.next
// 注意判断边界问题
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	cur := head
	curr := head

	r := make([]int, 0)
	for cur != nil {
		r = append(r, cur.Val)
		cur = cur.Next
	}

	idx := len(r) - n

	if idx >= len(r) {
		return head
	}

	if idx == 0 {
		return head.Next
	}
	for i := 1; i < idx; i++ {
		if i != idx {
			curr = curr.Next
		}
	}

	curr.Next = curr.Next.Next

	return head
}
