/*
 * @lc app=leetcode.cn id=1094 lang=golang
 *
 * [1094] 拼车
 *
 * https://leetcode-cn.com/problems/car-pooling/description/
 *
 * algorithms
 * Medium (55.85%)
 * Likes:    164
 * Dislikes: 0
 * Total Accepted:    43.8K
 * Total Submissions: 78.7K
 * Testcase Example:  '[[2,1,5],[3,3,7]]\n4'
 *
 * 车上最初有 capacity 个空座位。车 只能 向一个方向行驶（也就是说，不允许掉头或改变方向）
 *
 * 给定整数 capacity 和一个数组 trips ,  trip[i] = [numPassengersi, fromi, toi] 表示第 i
 * 次旅行有 numPassengersi 乘客，接他们和放他们的位置分别是 fromi 和 toi 。这些位置是从汽车的初始位置向东的公里数。
 *
 * 当且仅当你可以在所有给定的行程中接送所有乘客时，返回 true，否则请返回 false。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：trips = [[2,1,5],[3,3,7]], capacity = 4
 * 输出：false
 *
 *
 * 示例 2：
 *
 *
 * 输入：trips = [[2,1,5],[3,3,7]], capacity = 5
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= trips.length <= 1000
 * trips[i].length == 3
 * 1 <= numPassengersi <= 100
 * 0 <= fromi < toi <= 1000
 * 1 <= capacity <= 10^5
 *
 *
 */

// @lc code=start
func carPooling(trips [][]int, capacity int) bool {
	delta := make([]int, 1001)
	for _, v := range trips {
		delta[v[1]] += v[0]
		delta[v[2]] -= v[0]
	}
	cur := 0
	for _, v := range delta {
		cur += v
		if cur > capacity {
			return false
		}
	}
	return true

}

// @lc code=end

