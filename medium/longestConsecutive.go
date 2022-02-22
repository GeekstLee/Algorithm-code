package medium

/**
128. 最长连续序列
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

 

示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
 

提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-consecutive-sequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func longestConsecutive(nums []int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}
	numsMap := make(map[int]bool)
	//将数字存入map中
	for _, num := range nums {
		numsMap[num] = true
	}
	maxLen := 1
	//如果存在序列x,x+1,x+2,x+3，因此我们只要找到第一个数，跳过那些存在x-1,x-2...的数
	for _, num := range nums {
		_, ok := numsMap[num-1]
		if ok {
			continue
		}
		currentLen := 1 //num第一个
		for numsMap[num+1] {
			num++ //下一个数
			currentLen++
		}
		if currentLen > maxLen {
			maxLen = currentLen
		}

	}
	return maxLen
}
