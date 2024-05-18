package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		InvertTree(root.Left)
		InvertTree(root.Right)
		return root
	} else {
		return nil
	}
}
