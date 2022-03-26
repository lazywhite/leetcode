/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(n int) int {
	// n==0 --> 0
	// n==1 --> 1
	// n==2 --> 1
	// n==3 --> 2
	// n==4 --> 3
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n]
}

// @lc code=end

