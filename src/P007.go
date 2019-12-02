package src

import (
	"fmt"
	"strconv"
)

// 整数反转
// https://leetcode-cn.com/problems/reverse-integer/
func reverse(x int) int {
	xs := []byte(fmt.Sprint(x))
	sign := 1
	if xs[0] == '-' {
		xs = xs[1:]
		sign = -1
	}
	for i := 0; i < len(xs)/2; i++ {
		xs[i], xs[len(xs)-1-i] = xs[len(xs)-1-i], xs[i]
	}
	v, err := strconv.ParseInt(string(xs), 10, 32)
	if err != nil {
		return 0
	}
	return sign * int(v)
}
