package src

import "fmt"

// 回文数
// https://leetcode-cn.com/problems/palindrome-number/
func isPalindrome(x int) bool {
	xs := []byte(fmt.Sprint(x))
	for i := 0; i < len(xs)/2; i++ {
		if xs[i] != xs[len(xs)-1-i] {
			return false
		}
	}
	return true
}
