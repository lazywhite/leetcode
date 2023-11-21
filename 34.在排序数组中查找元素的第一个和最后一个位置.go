/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	empty := []int{-1, -1}
	if len(nums) == 0 {
		return empty
	}
	size := len(nums)
	findFirst := func() int {
		left := 0
		right := size - 1
		result := -1
		for left <= right {
			mid := left + (right-left)/2
			switch {
			case nums[mid] < target:
				left = mid + 1
			case nums[mid] == target:
				// 更新结果值, 继续向下搜索
				result = mid
				right = mid - 1
			case nums[mid] > target:
				right = mid - 1
			}
		}
		return result
	}
	findLast := func(left, right int) int {
		// 初始值设为right+1, 有可能[mid, right]全部为target值, 导致找不到比target大的
		// 因为最终要返回last - 1, 所以加一
		result := right + 1
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] <= target {
				left = mid + 1
			} else {
				result = mid
				right = mid - 1
			}
		}
		return result
	}
	first := findFirst()
	if first == -1 {
		return empty
	}
	last := findLast(first+1, size-1)
	return []int{first, last - 1}
}

// @lc code=end

