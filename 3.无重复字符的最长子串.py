#
# @lc app=leetcode.cn id=3 lang=python
#
# [3] 无重复字符的最长子串
#

# @lc code=start
class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        # 贪心算法+动态窗口(双指针)
        tmp = {}
        slow = 0
        fast = 0
        max_size = 0
        while fast < len(s):
            char = s[fast]
            if char in tmp:
                # 慢指针虽然会更新, 但tmp没有更新, 里面仍然包含已经被删掉的字符的下标
                # 因此必须取慢指针当前位置和缓存位置加一的最大值
                slow = max(slow, tmp[char] + 1)
            tmp[char] = fast
            max_size = max(max_size, fast - slow + 1)
            fast += 1
                
        return max_size


# @lc code=end

