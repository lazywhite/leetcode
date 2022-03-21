/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {

	r := x ^ y
	count := 0
	for i := 0; i < 32; i++ {
		flag := 1 << i
		if flag & r > 0{
			count++
		}
	}

	return count
}

// @lc code=end

