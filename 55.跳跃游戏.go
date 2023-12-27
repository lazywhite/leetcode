/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	/*
		贪心算法
		对于位置x, x + nums[x]是最远可达的位置, 意味着区间[x, x + nums[x]]内
		的所有位置都可以到
		用right_most记录上一次最远可达的位置, 如果本次x是大于这个位置的, 说明不可达, 直接返回false
	*/
	size := len(nums)
	if size <= 1 {
		return true
	}
	last := size - 1
	right_most := 0 // 不能是-1, 否则i <= pre_max 为false
	for i := 0; i <= last; i++ {
		if i <= right_most {
			// 本位置可达, 更新right_most
			cur := i + nums[i]
			right_most = max(right_most, cur)
			if right_most >= last {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

