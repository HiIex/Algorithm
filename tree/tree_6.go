package tree

/*
判断二叉树是否轴对称

转化为：
从根节点的左右孩子开始，递归判断节点A和B：
1. A和B的值是否相等
2. A的左孩子和B的右孩子是否对称
3. A的右孩子和B的左孩子是否对称
*/

func CheckSymmetricTree(root *TreeNode) bool {
	if root ==nil{
		return true
	}

	return compare2(root.Left,root.Right)
}

func compare2(a *TreeNode, b *TreeNode) bool{
   if a==nil && b==nil{
	   return true
   }

   if a!=nil && b==nil{
	   return false
   }

   if a==nil && b!=nil{
	   return false
   }

   if a.Val !=b.Val{
	   return false
   }

   return compare2(a.Left,b.Right) && compare2(a.Right,b.Left) 
}
