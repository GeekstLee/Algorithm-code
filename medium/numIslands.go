package medium

/**
200. 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

 

示例 1：

输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
示例 2：

输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3
 

提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	//岛屿数量
	ans := 0
	row, col := len(grid), len(grid[0])
	//标记属于该岛屿的小岛
	var dfs func(r, c int)
	dfs = func(r, c int) {
		//控制越界，排除不符合的岛屿(标记非1的)
		if r < 0 || r >= row || c < 0 || c >= col || grid[r][c] != '1' {
			return
		}
		//进行标记
		grid[r][c] = '2'
		dfs(r+1, c) //上
		dfs(r-1, c) //下
		dfs(r, c-1) //左
		dfs(r, c+1) //右

	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}
