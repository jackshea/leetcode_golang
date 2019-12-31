package src

import "math"

// 二叉树的最小深度
// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	minDep := math.MaxInt32
	if root.Left != nil {
		minDep = min(minDep, minDepth(root.Left))
	}

	if root.Right != nil {
		minDep = min(minDep, minDepth(root.Right))
	}

	return 1 + minDep
}

func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
