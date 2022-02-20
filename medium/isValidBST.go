package medium

import "math"

/**
98. 验证二叉搜索树

给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
 

示例 1：


输入：root = [2,1,3]
输出：true
示例 2：


输入：root = [5,1,4,null,null,3,6]
输出：false
解释：根节点的值是 5 ，但是右子节点的值是 4 。
 

提示：

树中节点数目范围在[1, 104] 内
-231 <= Node.val <= 231 - 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

二叉搜索树的概念：
（1）如果该二叉树的左子树不为空，则左子树上所有节点的值均小于它的根节点的值
（2）若它的右子树不空，则右子树上所有节点的值均大于它的根节点的值
（3）它的左右子树也为二叉搜索树

*/

func isValidBST(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	//对上下限进行限制，而不是左右节点
	var dfs func(node *TreeNode, upper int, lower int) bool
	dfs = func(node *TreeNode, upper int, lower int) bool {
		if node == nil {
			return true
		}
		//不合法的节点
		if node.Val <= lower || node.Val >= upper {
			return false
		}
		return dfs(node.Left, node.Val, lower) && dfs(node.Right, upper, node.Val)

	}
	return dfs(root, math.MaxInt64, math.MinInt64)
}

/**
中序遍历升序
*/

func isValidBSTM2(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	var nums = make([]int64, 0)
	var search func(node *TreeNode)
	search = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			search(node.Left)
		}
		nums = append(nums, int64(node.Val))
		if node.Right != nil {

			search(node.Right)
		}
	}
	search(root)
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			return false
		}
	}
	return true
}
