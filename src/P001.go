package src

// 两数之和
// https://leetcode-cn.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	vtoi := make(map[int]int)
	for k, v := range nums {
		vtoi[v] = k
	}

	for i := 0; i < len(nums); i++ {
		if index, ok := vtoi[target-nums[i]]; ok {
			if i != index {
				return []int{i, index}
			}
		}
	}
	return nil
}
