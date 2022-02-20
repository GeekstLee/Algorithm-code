package medium

/**
79. 单词搜索

给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

 

示例 1：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
输出：true
示例 3：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
输出：false
 

提示：

m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board 和 word 仅由大小写英文字母组成
 

进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func exist(board [][]byte, word string) bool {
	if board == nil || len(board) < 1 {
		return false
	}

	//string to byte
	str := []byte(word)
	row := len(board)
	col := len(board[0])
	if row*col < len(str) {
		return false
	}
	//初始化数组
	search := make([][]bool, len(board)) //标记是否被使用过
	for i := 0; i < len(search); i++ {
		search[i] = make([]bool, len(board[i]))
	}

	var dfs func(x int, y int, index int) bool
	dfs = func(x int, y int, index int) bool {
		if index == len(str) {
			return true
		}
		//不能越界，不能被访问过，字母要相等
		if x < 0 || x >= row || y < 0 || y >= col || search[x][y] || board[x][y] != str[index] {
			return false
		}

		//进行访问标记
		search[x][y] = true
		//从 上下左右方向 进行访问
		res1 := dfs(x, y-1, index+1) //向上
		res2 := dfs(x, y+1, index+1) //向下
		res3 := dfs(x-1, y, index+1) //向左
		res4 := dfs(x+1, y, index+1) //向右
		if res1 || res2 || res3 || res4 {
			return true
		}
		//取消标记
		search[x][y] = false
		return false
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
