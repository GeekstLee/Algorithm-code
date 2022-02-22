package medium

import (
	"math"
)

/**

221. 最大正方形
在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。



示例 1：


输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：4
示例 2：


输入：matrix = [["0","1"],["1","0"]]
输出：1
示例 3：

输入：matrix = [["0"]]
输出：0


提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 300
matrix[i][j] 为 '0' 或 '1'
*/
func maximalSquare(matrix [][]byte) int {
	if matrix == nil {
		return 0
	}
	maxLen := 0 //计算长度
	row, col := len(matrix), len(matrix[0])

	dp := make([][]int, row) //计算最大值
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxLen = 1
		}
	}
	for j := 0; j < col; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			maxLen = 1
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
			}
			if matrix[i-1][j] == '1' && matrix[i-1][j-1] == '1' && matrix[i][j-1] == '1' {
				//取最短，最长不一定能构成正方形
				dp[i][j] += int(math.Min(float64(dp[i-1][j]), math.Min(float64(dp[i-1][j-1]), float64(dp[i][j-1]))))
			}
			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
			}
		}
	}
	return maxLen * maxLen
}
