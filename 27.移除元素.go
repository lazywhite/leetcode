/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
// 双指针法, 将等于val的值全部移动到右面
func removeElement(nums []int, val int) int {

	// 存储第一个值为val的index
	slowIdx := 0

	for fastIdx := 0; fastIdx < len(nums); fastIdx++ {
		if nums[fastIdx] != val {
			if slowIdx != fastIdx {
				swap(nums, slowIdx, fastIdx)
			}
			// 可能连续多个item都为val, 因此只能加1
			slowIdx++
		}
	}
	nums = nums[:slowIdx]
	return slowIdx
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// @lc code=end

