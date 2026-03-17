package main

//是否是环形链表并返回相遇节点

// 哈希
//func detectCycle(head *ListNode) *ListNode {
//	m := make(map[*ListNode]bool)
//	for cur := head; cur != nil; cur = cur.Next {
//		if m[cur] {
//			return cur
//		}
//		m[cur] = true
//	}
//	return nil
//}

// 快慢指针
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			fast = head
			for fast != nil {
				if fast == slow {
					return slow
				}
				fast = fast.Next
				slow = slow.Next
			}

		}
	}
	return nil
}
