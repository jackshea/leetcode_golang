package src

// 搜索插入位置
// https://leetcode-cn.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			return mid
		} else {
			right = mid
		}
	}
	return left
}
