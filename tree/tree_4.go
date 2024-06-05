package tree

/*
判断B是否位A的子结构
A的每个节点都要遍历一遍，判断以该节点为根节点时，是否和B匹配
*/
func IsSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	//比较当前节点，递归比较左孩子节点和右孩子节点
	return compare(A, B) || IsSubStructure(A.Left, B) || IsSubStructure(A.Right, B)
}

// compare 匹配
//  @param a
//  @param b
//  @return bool
func compare(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return true
	}

	if a == nil || a.Val != b.Val {
		return false
	}

	//递归遍历两棵树的所有节点
	return compare(a.Left, b.Left) && compare(a.Right, b.Right)
}
