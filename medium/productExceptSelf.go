package medium

/**
238. 除自身以外数组的乘积
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请不要使用除法，且在 O(n) 时间复杂度内完成此题。



示例 1:

输入: nums = [1,2,3,4]
输出: [24,12,8,6]
示例 2:

输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]


提示：

2 <= nums.length <= 105
-30 <= nums[i] <= 30
保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内


进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
*/
func productExceptSelf(nums []int) []int {
	if nums == nil {
		return nil
	}
	L := len(nums)
	//前缀数组乘积
	prefix := make([]int, len(nums))
	prefix[0] = 1
	for i := 1; i < L; i++ {
		prefix[i] = nums[i-1] * prefix[i-1]
	}
	//后缀数组乘积
	suffix := make([]int, len(nums))
	suffix[L-1] = 1
	for i := L - 2; i >= 0; i-- {
		suffix[i] = nums[i+1] * suffix[i+1]
	}
	for i := 0; i < L; i++ {
		nums[i] = suffix[i] * prefix[i]
	}
	return nums
}
