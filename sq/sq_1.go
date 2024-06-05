package sq

/*
两个栈实现一个队列
一个正向栈，一个反向栈
*/

type CQueue struct {
	s1 []int
	s2 []int
}

func Constructor1() CQueue {
	return CQueue{
		s1: make([]int, 0),
		s2: make([]int, 0),
	}
}

// AppendTail 进队：进正向栈
//  @receiver q
//  @param value
func (q *CQueue) AppendTail(value int) {
	q.s1 = append(q.s1, value)
}

// DeleteHead 出队：考虑三种情况
//  @receiver q
//  @return int
func (q *CQueue) DeleteHead() int {
	if len(q.s2) == 0 && len(q.s1) == 0 {
		return -1
	}

	if len(q.s2) == 0 {
		for i := len(q.s1) - 1; i >= 1; i-- {
			q.s2 = append(q.s2, q.s1[i])
		}
		value := q.s1[0]
		q.s1 = make([]int, 0)
		return value
	}

	tail := len(q.s2) - 1
	value := q.s2[tail]
	q.s2 = q.s2[:tail]
	return value
}
