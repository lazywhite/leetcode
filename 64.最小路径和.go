/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 *
 * https://leetcode-cn.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (69.10%)
 * Likes:    1215
 * Dislikes: 0
 * Total Accepted:    349.2K
 * Total Submissions: 505.2K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 *
 * 说明：每次只能向下或者向右移动一步。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
 * 输出：7
 * 解释：因为路径 1→3→1→1→1 的总和最小。
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[1,2,3],[4,5,6]]
 * 输出：12
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1
 * 0
 *
 *
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	rows := len(grid)
	columns := len(grid[0])

	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
	}

	dp[0][0] = grid[0][0]

	// 初始化第一行
	for i := 1; i < columns; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	// 初始化第一列
	for j := 1; j < rows; j++ {
		dp[j][0] = dp[j-1][0] + grid[j][0]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
		}
	}
	return dp[rows-1][columns-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

