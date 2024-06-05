package tree

// DecorateRecord2 每一层分批输出：每次把队列内的所有元素一次性出队
//  @param root
//  @return []
func DecorateRecord2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	q := []*TreeNode{root}
	for {
		length := len(q)
		if length == 0 {
			return result
		}

		row := []int{}
		//一次性出队
		for i := 0; i < length; i++ {
			row = append(row, q[i].Val)
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		result = append(result, row)
		q = q[length:]
	}

}
