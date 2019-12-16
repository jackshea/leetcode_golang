package src

// 加一
// https://leetcode-cn.com/problems/plus-one/
func plusOne(digits []int) []int {
	add := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		add = digits[i] / 10
		digits[i] %= 10
		if add == 0 {
			break
		}
	}
	if add == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}
