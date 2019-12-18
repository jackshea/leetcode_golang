package src

// 爬楼梯
// https://leetcode-cn.com/problems/climbing-stairs/
func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}

	if n <= 2 {
		return n
	}

	a := 1
	b := 1
	c := a + b
	for i := 3; i <= n; i++ {
		a = b
		b = c
		c = a + b
	}
	return c
}
