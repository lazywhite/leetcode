/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
// 双指针法, 将等于val的值全部移动到右面
func removeElement(nums []int, val int) int {

	left, right := 0, 0

	for right < len(nums) {
		if nums[right] != val {
			if left == right {
				left++
			} else {
				swap(nums, left, right)
				left++
			}
		}
		right++
	}
	return left
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// @lc code=end

