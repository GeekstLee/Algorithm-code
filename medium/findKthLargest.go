package medium

import "fmt"

/**
215. 数组中的第K个最大元素
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

 

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
 

提示：

1 <= k <= nums.length <= 104
-104 <= nums[i] <= 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findKthLargest(nums []int, k int) int {
	left, right := 0, len(nums)-1
	target := len(nums) - k //目标位置
	index := -1
	for true {
		index = searchIndex(nums, left, right)
		if index == target {
			break
		}
		if index > target {
			right = index - 1
		} else {
			left = index + 1
		}
	}
	return nums[index]
}

//对[left..right]进行操作，返回 nums[left] 排序后的正确位置
//nums[left] 到 nums[j - 1] 中的所有元素都不大于 nums[j]     nums[j
//nums[j + 1] 到 nums[right] 中的所有元素都不小于 nums[j]
func searchIndex(nums []int, left int, right int) int {
	pivot := nums[left]
	j := left //存储nums[left]原本位置
	// 3 2 1 5 6 4
	fmt.Println(pivot, "===>", nums)
	for i := j + 1; i <= right; i++ {
		//nums[left+1...j] < pivot && nums[j+1...right] >= pivot
		if nums[i] < pivot {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	//对left和j进行交换，交换后的j位置就是pivot，j之前都是小于pivot，j之后大于等于pivot
	nums[j], nums[left] = nums[left], nums[j]
	return j
}
