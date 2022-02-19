package medium

/**
78. 子集
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

 

示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：

输入：nums = [0]
输出：[[],[0]]
 

提示：

1 <= nums.length <= 10
-10 <= nums[i] <= 10
nums 中的所有元素 互不相同

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func subsets(nums []int) [][]int {
	if len(nums) < 1 {
		return [][]int{{}}
	} else if len(nums) == 1 {
		return [][]int{{}, []int{nums[0]}}
	}
	ans := make([][]int, 0)
	ans = append(ans, []int{}) //先加入空集
	//每遍历到一个元素，先加入自身，再将前面所有子集加上当前元素成为新子集
	for i := 0; i < len(nums); i++ {
		for _, sun := range ans {
			//这里注意要进行copy操作
			copyRes := make([]int, len(sun))
			copy(copyRes, sun)
			ans = append(ans, append(copyRes, nums[i]))
		}

	}
	return ans
}
