package medium

import "sort"

/**
49. 字母异位词分组
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。

 

示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]
 

提示：

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/group-anagrams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func groupAnagrams(strs []string) [][]string {
	if strs == nil || len(strs) < 1 {
		return [][]string{}
	} else if len(strs) == 1 {
		return [][]string{strs}
	}
	sortMap := make(map[string][]string)
	var res [][]string
	for _, str := range strs {
		//对每个字符串进行排序，用排序后的子串作为key
		bStr := []byte(str)
		sort.Slice(bStr, func(i, j int) bool {
			return bStr[i] < bStr[j]
		})
		sortMap[string(bStr)] = append(sortMap[string(bStr)], str)
	}
	for _, str := range sortMap {
		res = append(res, str)
	}
	return res
}
