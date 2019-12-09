package src

// 实现 strStr()
// https://leetcode-cn.com/problems/implement-strstr/
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		isSame := true
		for j := 0; j < len(needle); j++ {
			if needle[j] != haystack[i+j] {
				isSame = false
				break
			}
		}
		if isSame {
			return i
		}
	}
	return -1
}
