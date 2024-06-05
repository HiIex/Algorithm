package tree

// DecorateRecor3 每一层分批输出，偶数层反序输出
//  @param root
//  @return []
func DecorateRecor3(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	q := []*TreeNode{root}
	reverse := false
	for {
		length := len(q)
		if length == 0 {
			return result
		}

		row := []int{}
		//每层一次性输出
		for i := 0; i < length; i++ {
			temp := []int{q[i].Val}
			//偶数层反序
			if reverse {
				row = append(temp, row...)
			} else {
				row = append(row, temp...)
			}

			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}

		result = append(result, row)
		q = q[length:]
		reverse = !reverse
	}

}
