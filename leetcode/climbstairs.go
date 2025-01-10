package leetcode

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	last := 1
	secondLast := 1
	final := 0

	for i := n - 1; i > 0; i-- {
		final = last + secondLast
		last = secondLast
		secondLast = final

	}
	return final
}
