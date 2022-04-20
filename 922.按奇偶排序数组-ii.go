/*
 * @lc app=leetcode.cn id=922 lang=golang
 *
 * [922] 按奇偶排序数组 II
 *
 * https://leetcode-cn.com/problems/sort-array-by-parity-ii/description/
 *
 * algorithms
 * Easy (71.32%)
 * Likes:    235
 * Dislikes: 0
 * Total Accepted:    96.2K
 * Total Submissions: 134.9K
 * Testcase Example:  '[4,2,5,7]'
 *
 * 给定一个非负整数数组 nums，  nums 中一半整数是 奇数 ，一半整数是 偶数 。
 *
 * 对数组进行排序，以便当 nums[i] 为奇数时，i 也是 奇数 ；当 nums[i] 为偶数时， i 也是 偶数 。
 *
 * 你可以返回 任何满足上述条件的数组作为答案 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,2,5,7]
 * 输出：[4,5,2,7]
 * 解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [2,3]
 * 输出：[2,3]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= nums.length <= 2 * 10^4
 * nums.length 是偶数
 * nums 中一半是偶数
 * 0 <= nums[i] <= 1000
 *
 *
 *
 *
 * 进阶：可以不使用额外空间解决问题吗？
 *
 */
package main

func main() {
	a := []int{2, 3}
	sortArrayByParityII(a)
}

// @lc code=start
func sortArrayByParityII(nums []int) []int {

	size := len(nums)
	left := 0
	right := size - 1

	for left <= size-2 && right >= 1 {
		for left <= size-2 && nums[left]%2 == 0 {
			left += 2
		}
		for right >= 1 && nums[right]%2 == 1 {
			right -= 2
		}
		if left < size && right >= 0{
			swap(nums, left, right)
			left += 2
			right -= 2

		}
	}
	return nums
}

func swap(arr []int, x, y int) {
	arr[x], arr[y] = arr[y], arr[x]
}

// @lc code=end
