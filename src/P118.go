package src

// 杨辉三角
// https://leetcode-cn.com/problems/pascals-triangle/
func generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}

	var ans [][]int
	lastRow := []int{1}
	ans = append(ans, lastRow)
	for i := 1; i < numRows; i++ {
		curRow := []int{1}
		for j := 1; j < len(lastRow); j++ {
			curRow = append(curRow, lastRow[j-1]+lastRow[j])
		}
		curRow = append(curRow, 1)
		lastRow = curRow
		ans = append(ans, lastRow)
	}
	return ans
}
