package src

// 最长公共前缀
// https://leetcode-cn.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := []byte(strs[0])
	for _, v := range strs {
		if len(v) < len(prefix) {
			prefix = prefix[:len(v)]
		}

		for i := 0; i < len(prefix); i++ {
			if prefix[i] != v[i] {
				prefix = prefix[:i]
			}
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return string(prefix)
}
