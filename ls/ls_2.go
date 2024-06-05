package ls

// DeleteNode 删除链表中指定值的节点（可能有多个），返回头节点
/*
1. 创建虚拟头节点
2. 遍历链表，维护当前节点和前节点：碰到对应值的节点连接前后节点
3. 返回虚拟头节点的后节点
*/
//  @param head 
//  @param val 
//  @return *ListNode 
func DeleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	preHead := &ListNode{
		Next: head,
	}
	pre := preHead
	//可以不要cur，加了可读性好一点
	cur := head
	for {
		if cur == nil {
			return preHead.Next
		}

		if cur.Val == val {
			pre.Next = cur.Next
			cur = cur.Next
			continue
		}

		pre = pre.Next
		cur = cur.Next
	}

}
