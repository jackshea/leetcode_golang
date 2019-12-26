package src

// 将有序数组转换为二叉搜索树
// https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	node := new(TreeNode)
	node.Val = nums[mid]
	node.Left = build(nums, left, mid-1)
	node.Right = build(nums, mid+1, right)
	return node
}
