/*
 * @lc app=leetcode.cn id=785 lang=golang
 *
 * [785] 判断二分图
 */

// @lc code=start
func isBipartite(graph [][]int) bool {
	/*
		通过DFS标记所有节点, 如果能恰好标记为两种颜色, 则属于二分图
		如果标记过程中发现相邻的节点已经被标记过, 且跟自己颜色相同, 则不是二分图
	*/
	size := len(graph)
	colors := make([]int, size)
	nocolor, red, green := 0, 1, 2
	var valid = true

	var dfs func(idx int, color int)
	dfs = func(idx, color int) {
		if !valid {
			return
		}
		colors[idx] = color // 记录节点被标记的颜色
		// 拿到相邻节点的颜色
		another := red
		if color == red {
			another = green
		}

		for _, i := range graph[idx] {
			// 没有被访问过, 标记为another
			if colors[i] == nocolor {
				dfs(i, another)
			}
			/// 已经被访问过, 但颜色相同
			if colors[i] == color {
				valid = false
				return
			}
			/// 已经被访问过, 颜色不同, 不做处理
		}
	}

	for i := 0; i < size; i++ {
		if colors[i] == nocolor {
			dfs(i, red) // 先标记为red
		}
	}
	return valid
}

// @lc code=end

