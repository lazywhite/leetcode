#
# @lc app=leetcode.cn id=739 lang=python
#
# [739] 每日温度
#

# @lc code=start
class Solution(object):
    def dailyTemperatures(self, temperatures):
        """
        :type temperatures: List[int]
        :rtype: List[int]
        """
        res = [None] * len(temperatures)
        q = []
        for i in range(len(temperatures)):
            while len(q) > 0 and temperatures[q[-1]] < temperatures[i]:
                idx = q.pop()
                res[idx] = i - idx
            q.append(i)
        for j in q:
            res[j] = 0
        return res

# @lc code=end

