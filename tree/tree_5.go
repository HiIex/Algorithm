package tree

/*
反转二叉树（镜像二叉树）：每个节点的左孩子和右孩子互换
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func MirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	mirror(root)
	MirrorTree(root.Left)
	MirrorTree(root.Right)

	return root
}

func mirror(root *TreeNode) {
	if root == nil {
		return
	}

	tempNode := root.Left
	root.Left = root.Right
	root.Right = tempNode
}
