package divideconquer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DeduceTree(preorder []int, inorder []int) *TreeNode {
	inIndexMap := make(map[int]int, 0)
	for index, val := range inorder {
		inIndexMap[val] = index
	}

	var reveal func(preRoot, inLeft, inRight int) *TreeNode
	reveal = func(preRoot, inLeft, inRight int) *TreeNode {
		if inLeft > inRight {
			return nil
		}

		rootVal := preorder[preRoot]
		inRoot := inIndexMap[rootVal]
		if inRoot < inLeft || inRoot > inRight {
			return nil
		}

		return &TreeNode{
			Val:   rootVal,
			Left:  reveal(preRoot+1, inLeft, inRoot-1),
			Right: reveal(preRoot+inRoot-inLeft+1, inRoot+1, inRight),
		}
	}

	return reveal(0, 0, len(inorder)-1)
}
