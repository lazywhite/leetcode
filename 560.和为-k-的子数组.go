/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 *
 * https://leetcode-cn.com/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (45.01%)
 * Likes:    1491
 * Dislikes: 0
 * Total Accepted:    226.9K
 * Total Submissions: 502.8K
 * Testcase Example:  '[1,1,1]\n2'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,1,1], k = 2
 * 输出：2
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,2,3], k = 3
 * 输出：2
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 2 * 10^4
 * -1000 <= nums[i] <= 1000
 * -10^7 <= k <= 10^7
 * 
 * 
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	pre, count := 0, 0
	m := make(map[int]int, 0)
	m[0] = 1
	for i := range nums {
		pre += nums[i]
		if v, ok := m[pre-k];ok{
			count += v
		}
		m[pre] += 1
	}
	return count
}
// @lc code=end

