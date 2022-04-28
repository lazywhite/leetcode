/*
 * @lc app=leetcode.cn id=1854 lang=golang
 *
 * [1854] 人口最多的年份
 *
 * https://leetcode-cn.com/problems/maximum-population-year/description/
 *
 * algorithms
 * Easy (76.07%)
 * Likes:    41
 * Dislikes: 0
 * Total Accepted:    13.9K
 * Total Submissions: 18.4K
 * Testcase Example:  '[[1993,1999],[2000,2010]]'
 *
 * 给你一个二维整数数组 logs ，其中每个 logs[i] = [birthi, deathi] 表示第 i 个人的出生和死亡年份。
 *
 * 年份 x 的 人口 定义为这一年期间活着的人的数目。第 i 个人被计入年份 x 的人口需要满足：x 在闭区间 [birthi, deathi - 1]
 * 内。注意，人不应当计入他们死亡当年的人口中。
 *
 * 返回 人口最多 且 最早 的年份。
 *
 *
 *
 * 示例 1：
 *
 * 输入：logs = [[1993,1999],[2000,2010]]
 * 输出：1993
 * 解释：人口最多为 1 ，而 1993 是人口为 1 的最早年份。
 *
 *
 * 示例 2：
 *
 * 输入：logs = [[1950,1961],[1960,1971],[1970,1981]]
 * 输出：1960
 * 解释：
 * 人口最多为 2 ，分别出现在 1960 和 1970 。
 * 其中最早年份是 1960 。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= logs.length <= 100
 * 1950 <= birthi < deathi <= 2050
 *
 *
 */

// @lc code=start
func maximumPopulation(logs [][]int) int {
	offset := 1950
	delta := make([]int, 101)
	for _, v := range logs {
		delta[v[0]-offset]++
		delta[v[1]-offset]--
	}
	max := 0
	idx := 0
	cur := 0
	for i := 1; i < 101; i++ {
		cur += delta[i]
		if cur > max {
			max = cur
			idx = i
		}
	}
	return idx + offset

}

// @lc code=end

