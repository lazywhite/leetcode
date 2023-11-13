#
# @lc app=leetcode.cn id=239 lang=python
#
# [239] 滑动窗口最大值
#

# @lc code=start
class Solution(object):
    def maxSlidingWindow(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: List[int]
        """
        deque = collections.deque()
        res = []
        for i, v in enumerate(nums):
            # 实现max deque, 顶部存储最大值
            # 1. 删除下标超出范围的
            if len(deque) > 0 and deque[0] <= i - k:
                deque.popleft()
            # 2. 删除比当前值小的所有
            while len(deque) > 0 and nums[deque[-1]] < v:
                deque.pop()
            deque.append(i)
            res.append(nums[deque[0]])
        return res[k-1:]
# @lc code=end

