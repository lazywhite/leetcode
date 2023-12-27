/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	/* 双指针法, 直接修改输入数组
	如果右指针指向的元素不等于val，它一定是输出数组的一个元素，我们就将右指针指向的元素复制到左指针位置，然后将左右指针同时右移；
	如果右指针指向的元素等于val，它不能在输出数组里，此时左指针不动，右指针右移一位。
	*/

	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}

// @lc code=end

