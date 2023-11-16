#
# @lc app=leetcode.cn id=463 lang=python3
#
# [463] 岛屿的周长
#

# @lc code=start
class Solution:
    def islandPerimeter(self, grid: List[List[int]]) -> int:
        rows = len(grid)
        if rows == 0:
            return 0
        cols = len(grid[0])
        def dfs(r, c):
            # 1. 超出边界
            if r < 0 or r >= rows or c < 0 or c >= cols:
                return 1
            # 2. 已被访问过
            if grid[r][c] == 2:
                return 0
            # 3. 属于水
            if grid[r][c] == 0:
                return 1
            grid[r][c] = 2 # 标记为已访问
            # 返回四个方向的sum
            return dfs(r-1, c) + dfs(r+1, c) + dfs(r, c-1) + dfs(r, c+1)
        res = 0
        for r in range(rows):
            for c in range(cols):
                if grid[r][c] == 1:
                    c = dfs(r, c)
                    res += c
        return res

# @lc code=end

