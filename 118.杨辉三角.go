/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {

	// 初始化最外层数组
	rt := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		rt[i] = make([]int, i+1) // 每一层初始化
		rt[i][0] = 1
		rt[i][i] = 1
		for j := 1; j < i; j++ {
			rt[i][j] = rt[i-1][j-1] + rt[i-1][j]
		}

	}
	return rt
}

// @lc code=end

