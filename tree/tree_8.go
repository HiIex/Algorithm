package tree

/*
计算二叉树的深度

重点：某个节点的深度=max(左子树的深度，右子树的深度)+1
*/

func CalculateDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := CalculateDepth(root.Left)
	rightDepth := CalculateDepth(root.Right)
	if leftDepth >= rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
