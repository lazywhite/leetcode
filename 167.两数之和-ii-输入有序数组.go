/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	// 双指针
	/*
		因为是有序数组, 则一定有
			nums[idx1-1] + nums[idx2] < target
			nums[idx1] + nums[idx2+1] > target
		因此肯定不会漏掉正确的
	*/

	i := 1
	j := len(numbers)

	result := make([]int, 2)

	for i < j {
		sum := numbers[i-1] + numbers[j-1]
		if sum == target {
			result[0] = i
			result[1] = j
			return result
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return result

}

// @lc code=end

