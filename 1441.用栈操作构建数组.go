/*
 * @lc app=leetcode.cn id=1441 lang=golang
 *
 * [1441] 用栈操作构建数组
 *
 * https://leetcode-cn.com/problems/build-an-array-with-stack-operations/description/
 *
 * algorithms
 * Easy (65.01%)
 * Likes:    40
 * Dislikes: 0
 * Total Accepted:    20.6K
 * Total Submissions: 31.7K
 * Testcase Example:  '[1,3]\n3'
 *
 * 给你一个目标数组 target 和一个整数 n。每次迭代，需要从  list = {1,2,3..., n} 中依序读取一个数字。
 *
 * 请使用下述操作来构建目标数组 target ：
 *
 *
 * Push：从 list 中读取一个新元素， 并将其推入数组中。
 * Pop：删除数组中的最后一个元素。
 * 如果目标数组构建完成，就停止读取更多元素。
 *
 *
 * 题目数据保证目标数组严格递增，并且只包含 1 到 n 之间的数字。
 *
 * 请返回构建目标数组所用的操作序列。
 *
 * 题目数据保证答案是唯一的。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：target = [1,3], n = 3
 * 输出：["Push","Push","Pop","Push"]
 * 解释：
 * 读取 1 并自动推入数组 -> [1]
 * 读取 2 并自动推入数组，然后删除它 -> [1]
 * 读取 3 并自动推入数组 -> [1,3]
 *
 *
 * 示例 2：
 *
 *
 * 输入：target = [1,2,3], n = 3
 * 输出：["Push","Push","Push"]
 *
 *
 * 示例 3：
 *
 *
 * 输入：target = [1,2], n = 4
 * 输出：["Push","Push"]
 * 解释：只需要读取前 2 个数字就可以停止。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= target.length <= 100
 * 1 <= target[i] <= 100
 * 1 <= n <= 100
 * target 是严格递增的
 *
 *
 */
// @lc code=start
func buildArray(target []int, n int) []string {
	ans := make([]string, 0)

	s := 1

	for i := 0; i < len(target); i++ {
		for target[i] != s {
			s++
			ans = append(ans, "Push", "Pop")
		}
		ans = append(ans, "Push")
		s++
	}

	return ans
}

// @lc code=end
