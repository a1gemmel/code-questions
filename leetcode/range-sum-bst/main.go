package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if low <= root.Val && root.Val <= high {
		return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	return rangeSumBST(root.Left, low, high)
}
