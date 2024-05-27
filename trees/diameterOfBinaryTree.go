package trees

// T - O(N) | SC - O(logN)
func DiameterOfBinaryTree(root *TreeNode) int {
	var (
		diameter int
		depth    func(node *TreeNode) int
	)

	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := depth(node.Left)
		right := depth(node.Right)

		if left+right > diameter {
			diameter = left + right
		}

		return max(left, right) + 1
	}

	depth(root)
	return diameter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
