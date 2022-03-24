/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 */

// @lc code=start
func findCircleNum(isConnected [][]int) int {

	/*
		1. isConnected[i][j] = 1  表示城市i与j相连
		2. 由于必须记载每个城市与其它城市的连接关系, 因此isConnected是一个n行n列的方阵
	*/

	//城市的数量
	cities := len(isConnected)

	//记录哪些城市被访问过
	visited := make([]bool, cities)

	// 访问与城市x相连的其它城市y1, y2..., 如果y1, y2没有被访问过, 则也进行访问
	var dfs func(int)
	dfs = func(x int) {
		visited[x] = true
		//cols表示x与其它n个城市的连接情况
		cols := isConnected[x]
		for y := 0; y < len(cols); y++ {
			// 如果y与city访问, 并且没有被访问过
			if cols[y] == 1 && !visited[y] {
				//查询城市y的访问情况
				dfs(y)
			}
		}
	}

	/*
	 1. 从城市0开始访问, 标记所有可达的城市, province=1
	 2. 如果所有的城市都被标记为true, 则province只有1个
	 3. 如果找到没有被访问过的城市, 证明存在第二个省, 调用dfs(), provice+=1
	 4. 最后所有的城市均被访问过
	*/
	province := 0
	for i, v := range visited {
		if v == false {
			province++
			dfs(i)
		}
	}
	return province

}

// @lc code=end

