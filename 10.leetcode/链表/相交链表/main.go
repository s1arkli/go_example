package main

/*
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			pa = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	tmp := make(map[*ListNode]bool)
	for v := headA; v != nil; v = v.Next {
		tmp[v] = true
	}
	for v := headB; v != nil; v = v.Next {
		if tmp[v] {
			return v
		}
	}
	return nil

}
