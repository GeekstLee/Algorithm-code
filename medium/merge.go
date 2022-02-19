package medium

import "sort"

/**
56. 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

 

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
 

提示：

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	var res [][]int
	//先排序，左区间从小到大
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var max func(a int, b int) int
	max = func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var min func(a int, b int) int
	min = func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < len(intervals); i++ {
		left, right := intervals[i][0], intervals[i][1] //获取当前集合的左右区间
		if len(res) == 0 || res[len(res)-1][1] < left {
			res = append(res, []int{left, right})
		} else {
			lastIndex := len(res) - 1
			//左区间取小，右区间取大
			res[lastIndex] = []int{min(res[lastIndex][0], left), max(res[lastIndex][1], right)}
		}
	}
	return res
}
