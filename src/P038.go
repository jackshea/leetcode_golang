package src

import (
	"fmt"
	"strings"
)

// 报数
// https://leetcode-cn.com/problems/count-and-say/
func countAndSay(n int) string {
	ans := "1"
	for i := 1; i < n; i++ {
		var sb strings.Builder
		var lastChar rune
		count := 0
		for k, v := range ans {
			if v == lastChar {
				count++
			} else {
				if k != 0 {
					sb.WriteString(fmt.Sprintf("%d%c", count, lastChar))
				}
				lastChar = v
				count = 1
			}
		}
		sb.WriteString(fmt.Sprintf("%d%c", count, lastChar))
		ans = sb.String()
	}
	return ans
}
