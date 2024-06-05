package tree

/*
判断是否为平衡二叉树（任意节点左子树的深度和右子树的深度差小于等于1）
*/

// 后序遍历：从下往上从左往右遍历每个节点，算节点的深度（算左子树、右子树的深度，顺便比较下差，差大于1就返回-1表示不平衡）
func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if depth2(root) != -1 {
		return true
	}

	return false
}

func depth2(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftDepth := depth2(node.Left)
	if leftDepth == -1 {
		return -1
	}

	rightDepth := depth2(node.Right)
	if rightDepth == -1 {
		return -1
	}

	delta := leftDepth - rightDepth
	if delta < -1 || delta > 1 {
		return -1
	}

	if leftDepth >= rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

// IsBalanced1 究极递归：层次遍历，每个节点算一遍左右子树的深度
//  @param root
//  @return bool
func IsBalanced1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	delta := depth(root.Left) - depth(root.Right)
	if delta < -1 || delta > 1 {
		return false
	}

	return IsBalanced1(root.Left) && IsBalanced1(root.Right)
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftDepth := depth(node.Left)
	rightDepth := depth(node.Right)
	if leftDepth >= rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
