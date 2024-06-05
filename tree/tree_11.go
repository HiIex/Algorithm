package tree

/*
寻找正常二叉树中两个节点的最邻近祖先
节点存在且值唯一
*/
func LowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p.Val == root.Val {
		return p
	}

	if q.Val == root.Val {
		return q
	}

	left := LowestCommonAncestor2(root.Left, p, q)
	right := LowestCommonAncestor2(root.Right, p, q)

	//左子树和右子树均不存在p或q
	if left == nil && right == nil {
		return nil
	}

	//右子树有p，q，p和q最邻近祖先三者的之一，并将该值往上一级节点传
	if left == nil {
		return right
	}

	//左子树有p，q，p和q最邻近祖先三者的之一，并将该值往上一级节点传
	if right == nil {
		return left
	}

	//p和q分处左右子树
	return root
}
