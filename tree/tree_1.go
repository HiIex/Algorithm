package tree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DecorateRecord 从左到右遍历二叉树（层次遍历），用队列就行
//  @param root
//  @return []int
func DecorateRecord(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	q := []*TreeNode{root}
	for {
		if len(q) == 0 {
			return result
		}

		result = append(result, q[0].Val)
		if q[0].Left != nil {
			q = append(q, q[0].Left)
		}
		if q[0].Right != nil {
			q = append(q, q[0].Right)
		}
		q = q[1:]
	}

}
