package src

import "math"

// 平衡二叉树
// https://leetcode-cn.com/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return depth(root) != -1
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftSubtreeDepth := depth(node.Left)
	if leftSubtreeDepth == -1 {
		return -1
	}
	rightSubtreeDepth := depth(node.Right)
	if rightSubtreeDepth == -1 {
		return -1
	}
	if math.Abs(float64(leftSubtreeDepth-rightSubtreeDepth)) > 1 {
		return -1
	}
	return 1 + max(leftSubtreeDepth, rightSubtreeDepth)
}

//func max(a int, b int) int {
//	if a >= b {
//		return a
//	} else {
//		return b
//	}
//}
