/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	/*
		贪心算法, 当发生重叠时, 优先删掉end比较大的区间
	*/
	// 按区间的start排序
	sort.Slice(intervals, func(x, y int) bool {
		return intervals[x][0] < intervals[y][0]
	})

	end := intervals[0][1]
	count := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			// 无重叠, 则更新end为当前区间的结束值
			end = intervals[i][1]
		} else {
			/* start与上一个end有重叠, 保留最小的end
			1. 如果current end < prev end, 说明删掉了上一个区间
			2. 如果current end > prev end, 说明删掉了现在区间
			*/
			end = min(end, intervals[i][1])
			count++
		}
	}
	return count
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

