package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxArr(arr []int) int {
	max := arr[0]
	for _, n := range arr {
		if n > max {
			max = n
		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -1 * a
}

func maxAncestorDiffInner(node *TreeNode, minVal, maxVal int) int {
	if node == nil {
		return 0
	}
	return maxArr([]int{
		abs(node.Val - minVal),
		abs(node.Val - maxVal),
		maxAncestorDiffInner(node.Left, min(minVal, node.Val), max(maxVal, node.Val)),
		maxAncestorDiffInner(node.Right, min(minVal, node.Val), max(maxVal, node.Val)),
	})
}

func maxAncestorDiff(root *TreeNode) int {
	return maxArr([]int{
		maxAncestorDiffInner(root.Left, root.Val, root.Val),
		maxAncestorDiffInner(root.Right, root.Val, root.Val),
	})
}
