package invertbinarytree

// TreeNode binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right, root.Left = root.Left, root.Right
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
