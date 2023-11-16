#
# @lc app=leetcode.cn id=39 lang=python
#
# [39] 组合总和
#

# @lc code=start
class Solution(object):
    def combinationSum(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        ans = []
        def dfs(target, combine, idx):
            if idx == len(candidates):
                return
            if target < 0:
                return
            if target == 0:
                ans.append(combine)
                return
            dfs(target, combine, idx+1) 
            dfs(target - candidates[idx], combine + [candidates[idx]], idx) # 不能使用combine.append, 必须新生成一个list

        dfs(target, [], 0)
        return ans
# @lc code=end

