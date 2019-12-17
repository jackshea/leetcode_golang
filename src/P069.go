package src

import "math"

//  x 的平方根
// https://leetcode-cn.com/problems/sqrtx/
func mySqrt(x int) int {
	left := 0
	right := int(math.Min(math.Pow(2, 16), float64(x)))
	for left < right {
		mid := left + (right-left+1)/2
		sqr := mid * mid
		if sqr == x {
			return mid
		} else if sqr < x {
			left = mid
		} else if sqr > x {
			right = mid - 1
		}
	}
	return left
}
