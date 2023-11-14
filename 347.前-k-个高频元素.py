#
# @lc app=leetcode.cn id=347 lang=python
#
# [347] 前 K 个高频元素
#

# @lc code=start
class Solution(object):
    def topKFrequent(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: List[int]
        """
        tmp = {}
        for i in nums:
            if i in tmp:
                tmp[i]+= 1
            else:
                tmp[i] = 1
        q = [(val, key) for key, val in tmp.items()]
        return [e[1] for e in heapq.nlargest(k, q)]
# @lc code=end

