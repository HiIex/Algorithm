package ls

// CopyRandomList 复杂链表深拷贝
//  @param head
//  @return *ListNode
func CopyRandomList(head *ListNode) *ListNode {
    //先遍历一遍在map中确定各个原节点与新节点的对应位置
	cur := head
	nodeMap := make(map[*ListNode]*ListNode, 0)
	len := 0
	for {
		if cur == nil {
			break
		}

		temp := &ListNode{
			Val: cur.Val,
		}
		nodeMap[cur] = temp

		cur = cur.Next
		len++
	}

    //再遍历一遍原链表，按原链表中的next关系与random关系与新链表中一一对应
	cur = head
	for i := 0; i < len; i++ {
		next := cur.Next
		random := cur.Random

		if next != nil {
			nodeMap[cur].Next = nodeMap[next]
		} else {
			nodeMap[cur].Next = nil
		}

		if random != nil {
			nodeMap[cur].Random = nodeMap[random]
		} else {
			nodeMap[cur].Random = nil
		}

		cur = cur.Next
	}

	return nodeMap[head]
}
