package ls

// GetIntersectionNode 找第一个交叉点（后续都一样），双向奔赴法
//  @param headA 
//  @param headB 
//  @return *ListNode 
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA==nil{
        return nil
    }

    if headB==nil{
        return nil
    }

    tempA:=headA
    tempB:=headB
    for {
        if tempA==tempB{
            return tempA
        }

        if tempA==nil{
            tempA=headB
        }else{
            tempA=tempA.Next
        }

        if tempB==nil{
            tempB=headA
        }else{
            tempB=tempB.Next
        }
    }
}