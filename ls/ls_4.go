package ls

// CountBackwards  取链表中倒数第cnt个节点的值
// 快慢指针
//  @param head
//  @param cnt
//  @return *ListNode
func CountBackwards(head *ListNode, cnt int) *ListNode {
	fast := head
	slow := head

	for i := 0; i < cnt; i++ {
		fast = fast.Next
	}

	for {
		if fast == nil {
			return slow
		}

		fast = fast.Next
		slow = slow.Next
	}
}
