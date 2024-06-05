package sq

type MinStack struct {
    s1 []int
    s2 []int 
}

/* 
最小栈：实现常数时间复杂度返回栈内最小元素

维护一个非严格降序辅助栈
*/



func Constructor2() MinStack {
    return MinStack{
        s1:make([]int,0),
        s2:make([]int,0),
    }
}


func (minStack *MinStack) Push(x int)  {
    minStack.s1=append(minStack.s1,x)
	// ！！！：<=很重要！，因为两个相同的最小元素pop掉一个后，在辅助栈中得留一个一样的
    if len(minStack.s2)==0 || x<=minStack.s2[len(minStack.s2)-1]{
        minStack.s2=append(minStack.s2,x)
    }
}


func (minStack *MinStack) Pop()  {
    top1:=len(minStack.s1)-1
    topValue:=minStack.s1[top1]
    minStack.s1=minStack.s1[:top1]

    top2:=len(minStack.s2)-1 
    min:=minStack.s2[top2]
    if topValue==min{
        minStack.s2=minStack.s2[:top2]
    }
}


func (minStack *MinStack) Top() int {
    return minStack.s1[len(minStack.s1)-1]
}


func (minStack *MinStack) GetMin() int {
    return minStack.s2[len(minStack.s2)-1]
}