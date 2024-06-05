package ls

import (
	"fmt"
)

// ReverseList 递归
//
//	@param head
//	@return *ListNode
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	next := head.Next
	head.Next = nil
	return reverser(head, next)

}

func reverser(reversedHead *ListNode, todoHead *ListNode) *ListNode {
	if todoHead == nil {
		return reversedHead
	}

	nextTodoHead := todoHead.Next
	todoHead.Next = reversedHead
	fmt.Printf("reversed head: %d, todo head: %d\n", reversedHead.Val, todoHead.Val)
	return reverser(todoHead, nextTodoHead)
}

// ReverseList2 头插
//
//	@param head
//	@return *ListNode
func ReverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	pre := head
	cur := head.Next
	head.Next = nil
	for {
		if cur == nil {
			return pre
		}

		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next

	}
}
