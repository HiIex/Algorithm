package tree

/*
返回二叉树中第cnt大的值

中序遍历时左孩子和右孩子的顺序倒换变成从大到小的顺序
*/

var count int
var val int

func FindTargetNode(root *TreeNode, cnt int) int {
	count = 0
	dfs(root, cnt)
	return val
}

func dfs(node *TreeNode, cnt int) {
	if count > cnt {
		return
	}

	if node.Right != nil {
		dfs(node.Right, cnt)
	}

	count++
	if count == cnt {
		val = node.Val
		return
	}

	if node.Left != nil {
		dfs(node.Left, cnt)
	}
}
