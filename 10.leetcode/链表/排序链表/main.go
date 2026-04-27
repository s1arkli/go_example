package main

/*
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head

	mid := findMiddle(head)
	rightHead := mid.Next
	mid.Next = nil

	right := sortList(rightHead)
	left := sortList(curr)

	return mergeList(right, left)
}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}
