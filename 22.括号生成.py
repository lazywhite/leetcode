#
# @lc app=leetcode.cn id=22 lang=python
#
# [22] 括号生成
#

# @lc code=start
class Solution(object):
    def generateParenthesis(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        res = []
        def backtrack(prefix, left, right):
            if left > n or right > left:
                return
            if len(prefix) == n * 2:
                res.append(prefix)
            backtrack(prefix + '(', left + 1, right)
            backtrack(prefix + ')', left, right + 1)
        backtrack("", 0, 0)
        return res
            
# @lc code=end

