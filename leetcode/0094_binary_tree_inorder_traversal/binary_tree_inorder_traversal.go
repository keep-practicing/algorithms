package binarytreeinordertraversal

// TreeNode binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 1)
	inorderTraversalHelp(root, &res)
	return res
}

func inorderTraversalHelp(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorderTraversalHelp(root.Left, res)
	*res = append(*res, root.Val)
	inorderTraversalHelp(root.Right, res)
}
