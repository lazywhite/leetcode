/*
 * @lc app=leetcode.cn id=1615 lang=golang
 *
 * [1615] 最大网络秩
 */

// @lc code=start
func maximalNetworkRank(n int, roads [][]int) int {
	// n*n矩阵, 保存城市间的连接信息
	conn := make([][]int, n)
	for i := 0; i < n; i++ {
		conn[i] = make([]int, n)
	}
	degree := make([]int, n) // 每个城市的度
	for _, v := range roads {
		conn[v[0]][v[1]] = 1
		conn[v[1]][v[0]] = 1
		degree[v[0]]++
		degree[v[1]]++
	}
	maxRank := -1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			/*
				题目并没有要求必须是两个直接相连的城市的rank
			*/
			rank := degree[i] + degree[j]
			if conn[i][j] == 1 {
				//如果相连, 路被多算了一次, 需要减掉
				rank--
			}
			maxRank = max(maxRank, rank)
		}
	}
	return maxRank
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

