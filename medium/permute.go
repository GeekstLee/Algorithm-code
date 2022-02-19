package medium

/**
46. 全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

 

示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]
 

提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func permute(nums []int) [][]int {
	if len(nums) < 1 {
		return [][]int{nums}
	}
	ans := make([][]int, 0)

	var dfs func(path []int)

	dfs = func(path []int) {
		if len(nums) == 0 {
			copyRes := make([]int, len(path))
			copy(copyRes, path)
			ans = append(ans, copyRes)
			return
		}
		for i, num := range nums {
			path = append(path, num)
			nums = append(nums[:i], nums[i+1:]...) //删除当前元素
			dfs(path)
			nums = append(nums[:i], append([]int{num}, nums[i:]...)...) //还原删除元素
			path = path[:len(path)-1]
		}
	}
	dfs([]int{})
	return ans
}
