package main

func main() {
	removeDuplicates([]int{1, 2, 2, 3, 3, 4})

}

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}
	left := 1
	for right := 1; right < size; right++ {
		if nums[right] != nums[right-1] {
			// 直接用right值覆盖left
			nums[left] = nums[right]
			left++
		}
	}
	return left

}
func swap(nums []int, x, y int) {
	nums[x], nums[y] = nums[y], nums[x]
}

// @lc code=end

// 冒泡移动到最后
func bak(nums []int) int {

	size := len(nums)

	// 从第2个开始向前比较
	i := 1
	for i < size {
		if nums[i] == nums[i-1] {
			// move to last
			for j := i; j < size-1; j++ {
				swap(nums, j, j+1)
			}
			// exclude last
			size--
		}
		i++
	}
	return size

}
