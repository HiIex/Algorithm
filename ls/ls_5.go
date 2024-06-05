package ls

// MergeOrderedList 按顺序合并两个有序的链表
// 
//  @param l1 
//  @param l2 
//  @return *ListNode 
func MergeOrderedList(l1 *ListNode, l2 *ListNode) *ListNode{
	if l1==nil{
        return l2
    }

    if l2==nil{
       return l1 
    }
    
    head:=&ListNode{}
    cur:=head
    for{
        if l1==nil{
            cur.Next=l2
            return head.Next
        }

        if l2==nil{
            cur.Next=l1
            return head.Next
        }

        if l1.Val<=l2.Val{
            cur.Next=l1
            l1=l1.Next
        }else{
            cur.Next=l2
            l2=l2.Next
        }

        cur=cur.Next
    }
}