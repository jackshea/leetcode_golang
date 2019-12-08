package src

// 移除元素
// https://leetcode-cn.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	curIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[curIndex] = nums[i]
			curIndex++
		}
	}

	return curIndex
}
