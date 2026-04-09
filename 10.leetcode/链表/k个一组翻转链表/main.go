package main

/*
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}
	curr := dummy
	end := dummy

	for {
		//判断是否符合反转条件
		for i := 1; i <= k; i++ {
			end = end.Next
			if end == nil {
				return dummy.Next
			}
		}

		//缓存下一个节点，可能是nil
		next := end.Next
		newHead := curr.Next

		//打断链表进行翻转
		end.Next = nil
		reverse(curr, curr.Next)

		//接上后续
		newHead.Next = next
		curr = newHead
		end = newHead
	}

}

func reverse(prev, a *ListNode) *ListNode {
	curr := a
	var subPrev *ListNode

	for curr != nil {
		next := curr.Next
		curr.Next = subPrev
		subPrev = curr
		curr = next
	}
	prev.Next = subPrev
	return subPrev
}
