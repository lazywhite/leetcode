/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		cap := (right - left) * min(height[left], height[right])
		if cap > max {
			max = cap
		}
		// 随着指针靠拢, 只有两个边界的最小值超过之前的最小值, 面积才可能变大
		// 因此即使两个边界值相等, 移动任何一个都是可以的, 不会漏掉最大值
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

