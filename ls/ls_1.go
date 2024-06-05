package ls


// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
	Random *ListNode
}

// ReverseBookList 反转链表到数组
/*
1. 先遍历一遍链表得到长度创建对应长度的数组，从数组末尾挨个再遍历一遍
2. 创建对应长度的数组
3. 再遍历一遍链表，从数组末尾开始往前填充值
*/
//  @param head
//  @return []int
func ReverseBookList(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	len := 1
	temp := head
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
		len++
	}

	array := make([]int, len)
	for i := len - 1; i >= 0; i-- {
		array[i] = head.Val
		head = head.Next
	}

	return array
}


