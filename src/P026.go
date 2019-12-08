package src

// 删除排序数组中的重复项
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	curIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[curIndex] != nums[i] {
			curIndex++
			nums[curIndex] = nums[i]
		}
	}
	return curIndex + 1
}
