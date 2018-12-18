package minimumdepthofbinarytree

import "algorithms/utils"

// TreeNode binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if left == 0 || right == 0 {
		return utils.CalcMaxInt(left, right) + 1
	}
	return utils.CalcMinInt(left, right) + 1
}
