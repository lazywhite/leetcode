/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {

	idx2 := n - 1
	idx1 := m - 1
	tail := m + n - 1

	/*
		如果使用归并排序, 需要额外的m+n长度数组, 空间复杂度较高
		nums1长度已经是m+n了, 后半部分是空的
		逆向双指针法, 可以保证nums1元素不会被覆盖
	*/
	for idx1 >= 0 && idx2 >= 0 {
		if nums2[idx2] >= nums1[idx1] {
			nums1[tail] = nums2[idx2]
			idx2--
		} else {
			nums1[tail] = nums1[idx1]
			idx1--
		}
		tail--
	}
	// 某个队列已耗尽
	if idx1 < 0 && idx2 >= 0 {
		for idx2 >= 0 {
			nums1[tail] = nums2[idx2]
			idx2--
			tail--
		}
	}
	if idx2 < 0 && idx1 >= 0 {
		for idx1 >= 0 {
			nums1[tail] = nums1[idx1]
			idx1--
			tail--
		}
	}

}

// @lc code=end

