package medium

/**
22. 括号生成
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

 

示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]
 

提示：

1 <= n <= 8

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
var result []string

func generateParenthesis(n int) []string {
	if n < 1 {
		return []string{}
	}
	result = []string{}
	backTrack("", n, n)
	return result
}

/**
使用回溯法
控制剩余的左括号和右括号数量
*/
func backTrack(str string, leftNum int, rightNum int) {
	if rightNum < leftNum { //不合法
		return
	}
	if leftNum == 0 && rightNum == 0 { //记录一个结果
		result = append(result, str)
		return
	}
	//尝试左括号
	if leftNum > 0 {
		backTrack(str+"(", leftNum-1, rightNum)
	}
	//尝试右括号
	if rightNum > 0 {
		backTrack(str+")", leftNum, rightNum-1)
	}
}
