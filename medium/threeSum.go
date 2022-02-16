package medium

import (
	"sort"
)

/**
15. 三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

 

示例 1：
-4 -1 -1 0 1 2
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：

输入：nums = []
输出：[]
示例 3：

输入：nums = [0]
输出：[]
 

提示：

0 <= nums.length <= 3000
-105 <= nums[i] <= 105

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	//排序，使得结果唯一
	sort.Ints(nums)
	res := make([][]int, 0)

	//枚举《第一个》元素
	for first := 0; first < len(nums); first++ {
		if nums[first] > 0 { //第一个数大于0，那后面怎么组都不可能合适，则退出
			break
		}
		//确保第一轮元素唯一，不能重复
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		target := 0 - nums[first] //剩余两个元素目标之和
		//枚举《第二个》元素
		for second := first + 1; second < len(nums); second++ {
			//确保元素唯一
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			//第三个元素，右指针
			third := len(nums) - 1
			for ; second < third && nums[second]+nums[third] > target; {
				third--
			}
			if third == second {
				break
			}
			if nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}

	}
	return res
}
