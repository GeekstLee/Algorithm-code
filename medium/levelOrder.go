package medium

/**
102. 二叉树的层序遍历
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

 

示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]
 

提示：

树中节点数目在范围 [0, 2000] 内
-1000 <= Node.val <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := make([][]int, 0)
	floor := make([]*TreeNode, 0) //上层节点
	floor = append(floor, root)
	for true {
		nums := make([]int, 0)
		nextFloor := make([]*TreeNode, 0)
		if len(floor) == 0 {
			break
		}
		for i := 0; i < len(floor); i++ {
			node := floor[i]
			nums = append(nums, node.Val)
			if node.Left != nil {
				nextFloor = append(nextFloor, node.Left)
			}
			if node.Right != nil {
				nextFloor = append(nextFloor, node.Right)
			}
		}

		ans = append(ans, nums)
		floor = make([]*TreeNode, len(nextFloor))
		copy(floor, nextFloor)
	}
	return ans
}
