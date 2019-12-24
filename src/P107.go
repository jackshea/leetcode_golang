package src

// 二叉树的层次遍历 II
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var opened []*TreeNode
	opened = append(opened, root)
	var ans [][]int
	for len(opened) != 0 {
		len := len(opened)
		var line []int
		for i := 0; i < len; i++ {
			node := opened[i]
			line = append(line, opened[i].Val)
			if node.Left != nil {
				opened = append(opened, node.Left)
			}

			if node.Right != nil {
				opened = append(opened, node.Right)
			}
		}
		opened = opened[len:]
		ans = append([][]int{line}, ans...)
	}
	return ans
}
