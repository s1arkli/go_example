package main

/*
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		a, b := prev.Next, prev.Next.Next

		prev.Next = b
		a.Next = b.Next
		b.Next = a

		prev = a
	}
	return dummy.Next
}
