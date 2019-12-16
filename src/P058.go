package src

// 最后一个单词的长度
// https://leetcode-cn.com/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
	ans := 0
	trimPreBlank := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			trimPreBlank = true
			ans++
		}

		if trimPreBlank && s[i] == ' ' {
			break
		}
	}
	return ans
}
