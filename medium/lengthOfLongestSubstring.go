package medium

/**
3. 无重复字符的最长子串
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

 

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
示例 4:

输入: s = ""
输出: 0
 

提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s) - 0
	}
	res := 0
	//字母表
	letterMap := make(map[string]int, len(s))
	left, right := 0, 0 //左右指针
	for ; right < len(s); right++ {
		var letter string
		//当前位置的字母
		if right == len(s)-1 { //最后一个字符
			letter = s[right:]
		} else {
			letter = s[right : right+1]
		}

		_, ok := letterMap[letter]
		letterMap[letter] += 1

		size := len(letterMap)
		if size > res {
			res = size
		}

		if !ok {
			continue
		}
		//缩小左边边界，直到字母不重复
		for {
			if letterMap[letter] <= 1 && left < len(s) {
				break
			}
			var str string
			if left == len(s)-1 {
				str = s[left:]
			} else {
				str = s[left : left+1]
			}
			if str == letter {
				letterMap[letter] -= 1
			} else {
				delete(letterMap, str)
			}
			left++
		}
	}
	return res
}
