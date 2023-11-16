#
# @lc app=leetcode.cn id=200 lang=python
#
# [200] 岛屿数量
#

# @lc code=start
class Solution(object):
    def numIslands(self, grid):
        """
        :type grid: List[List[str]]
        :rtype: int
        """
        # grid视为无向图, 格子为点, 相邻的格子间有边
        # 使用dfs进行处理
        rows = len(grid)
        if rows == 0:
            return 0
        cols = len(grid[0])
        def dfs(r, c):
            '''
            作用是将相邻的1置为0
            '''
            grid[r][c] = "0"
            for x, y in [(r-1, c), (r+1, c), (r, c-1), (r, c+1)]:
                if 0 <= x < rows and 0 <= y < cols and grid[x][y] == "1":
                    dfs(x, y)

        count = 0
        for r in range(rows):
            for c in range(cols):
                if grid[r][c] == "1":
                    count += 1
                    dfs(r, c)
        return count

        
# @lc code=end

