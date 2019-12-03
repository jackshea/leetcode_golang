package src

// 罗马数字转整数
// https://leetcode-cn.com/problems/roman-to-integer/
func romanToInt(s string) int {
	dic := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000}

	var last byte
	ans := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		ans += dic[c]
		if last == 'I' {
			if c == 'V' || c == 'X' {
				ans -= dic[last] * 2
			}
		} else if last == 'X' {
			if c == 'L' || c == 'C' {
				ans -= dic[last] * 2
			}
		} else if last == 'C' {
			if c == 'D' || c == 'M' {
				ans -= dic[last] * 2
			}
		}
		last = c
	}

	return ans
}
