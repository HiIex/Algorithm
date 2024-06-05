package tree

/*
二叉排序树：确认两个节点的最邻近祖先
节点存在且值唯一
*/

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//能进入到这里，说明p和q肯定同在以root为根节点的树上
	if root == nil {
		return nil
	}

	if p.Val == root.Val {
		return p
	}

	if q.Val == root.Val {
		return q
	}

	//同时小于根节点，说明p和q同处于左子树
	if p.Val < root.Val && q.Val > root.Val {
		return LowestCommonAncestor(root.Left, p, q)
	}

	//同时大于根节点，说明p和q同处于右子树
	if p.Val > root.Val && q.Val > root.Val {
		return LowestCommonAncestor(root.Right, p, q)
	}

	//一大一小，说明p和q分处不同子树
	return root
}
