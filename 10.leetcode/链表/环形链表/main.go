package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//哈希解法，暴力解法
//func hasCycle(head *ListNode) bool {
//	m := make(map[*ListNode]bool)
//	for cur := &head; cur != nil; cur = cur.Next {
//
//		if m[cur] {
//			return true
//		}
//		m[cur] = true
//	}
//
//	return false
//}

// 快慢指针
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
