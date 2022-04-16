/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 *
 * https://leetcode-cn.com/problems/insert-interval/description/
 *
 * algorithms
 * Medium (41.41%)
 * Likes:    569
 * Dislikes: 0
 * Total Accepted:    105.3K
 * Total Submissions: 254.1K
 * Testcase Example:  '[[1,3],[6,9]]\n[2,5]'
 *
 * 给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
 *
 * 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
 * 输出：[[1,5],[6,9]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
 * 输出：[[1,2],[3,10],[12,16]]
 * 解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
 *
 * 示例 3：
 *
 *
 * 输入：intervals = [], newInterval = [5,7]
 * 输出：[[5,7]]
 *
 *
 * 示例 4：
 *
 *
 * 输入：intervals = [[1,5]], newInterval = [2,3]
 * 输出：[[1,5]]
 *
 *
 * 示例 5：
 *
 *
 * 输入：intervals = [[1,5]], newInterval = [2,7]
 * 输出：[[1,7]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * intervals[i].length == 2
 * 0
 * intervals 根据 intervals[i][0] 按 升序 排列
 * newInterval.length == 2
 * 0
 *
 *
 */

// @lc code=start
func insert(input [][]int, newInterval []int) (result [][]int) {
	newStart := newInterval[0]
	newEnd := newInterval[1]

	merged := false
	for _, v := range input {
		start := v[0]
		end := v[1]

		if end < newStart {
			// 此区间在新区间左侧
			result = append(result, v)
		} else if start > newEnd {
			// 此区间在新区间右侧
			if !merged {
				// 检查新区间是否已经被插入过了
				result = append(result, []int{newStart, newEnd})
				merged = true
			}
			// 本区间一定需要插入
			result = append(result, v)
		} else {
			// 肯定相交了, 与当前区间进行合并
			newStart = min(start, newStart)
			newEnd = max(end, newEnd)
		}
	}
	if !merged {
		result = append(result, []int{newStart, newEnd})
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

