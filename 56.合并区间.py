#
# @lc app=leetcode.cn id=56 lang=python
#
# [56] 合并区间
#

# @lc code=start
class Solution(object):
    def merge(self, intervals):
        """
        :type intervals: List[List[int]]
        :rtype: List[List[int]]
        """
        # 先排序
        intervals.sort()
        merged = []
        for iv in intervals:
            if not merged or merged[-1][1] < iv[0]:
                merged.append(iv)
            else:
                merged[-1][1] = max(merged[-1][1], iv[1])
        return merged
# @lc code=end

