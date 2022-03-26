/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {

	cache := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if v, ok := cache[nums[i]]; ok {
			if abs(v, i) <= k {
				return true
			}

			// 将最大的index放入cache
			cache[nums[i]] = i

		} else {
			cache[nums[i]] = i
		}
	}
	return false
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

// @lc code=end

