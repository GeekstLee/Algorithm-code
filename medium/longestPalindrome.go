package medium

/**
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。

 

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
 

提示：

1 <= s.length <= 1000
s 仅由数字和英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	maxLen := 1                  //最大长度
	begin := 0                   //起始位置
	dp := make([][]bool, len(s)) //使用dp，dp[i][j]代表s中 i..j 字串是否为回文串
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ { //初始化：单个字符都是回文
		dp[i][i] = true
	}
	strArr := []byte(s)
	for Len := 2; Len <= len(s); Len++ { //枚举字串长度
		for left := 0; left < len(s); left++ { //枚举字串左边界
			right := left + Len - 1 //右边界
			if right >= len(s) {    //越界
				break
			}
			if strArr[left] != strArr[right] {
				dp[left][right] = false
			} else {
				if right-left < 3 { //对于中间只有一个字符的，肯定是回文串
					dp[left][right] = true
				} else { //中间字符串是回文串，则最终整个字符串都是回文串
					dp[left][right] = dp[left+1][right-1]
				}
			}
			if dp[left][right] && maxLen < (right-left+1) { //计算结果
				maxLen = right - left + 1
				begin = left
			}
		}
	}
	return s[begin : begin+maxLen]
}

/**
中心扩散法
*/
func longestPalindromeM2(s string) string {
	if len(s) < 2 {
		return s
	}
	begin, end := 0, 0 //位置下标
	for i := 0; i < len(s); i++ {
		//中心分两种：一种是一个点开始扩散【 a "b" a】，一种是两个点为中心扩散【a "bb" a】
		l1, r1 := expandByCenter([]byte(s), i, i)
		l2, r2 := expandByCenter([]byte(s), i, i+1)
		if r1-l1 > end-begin {
			begin = l1
			end = r1
		}
		if r2-l2 > end-begin {
			begin = l2
			end = r2
		}
	}
	return s[begin : end+1]
}

func expandByCenter(strArr []byte, left int, right int) (int, int) {
	for ; left >= 0 && right < len(strArr) && strArr[left] == strArr[right]; {
		left--
		right++
	}
	return left + 1, right - 1
}
