package src

// 路径总和
// https://leetcode-cn.com/problems/path-sum/
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	} else if hasPathSum(root.Left, sum-root.Val) {
		return true
	} else if hasPathSum(root.Right, sum-root.Val) {
		return true
	} else {
		return false
	}
}
