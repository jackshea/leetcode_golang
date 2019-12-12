package src

// 最大子序和
// https://leetcode-cn.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	ans := nums[0]
	sum := 0
	for _, v := range nums {
		sum += v
		if sum > ans {
			ans = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	return ans
}
