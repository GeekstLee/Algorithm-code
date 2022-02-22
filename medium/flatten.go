package medium

/**
114. 二叉树展开为链表
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。
 

示例 1：


输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [0]
输出：[0]
 

提示：

树中结点数在范围 [0, 2000] 内
-100 <= Node.val <= 100
 

进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	//获得前序遍历得到的节点
	nodes := preorder(root)
	for i := 1; i < len(nodes); i++ {
		pre, now := nodes[i-1], nodes[i]
		pre.Left = nil
		pre.Right = now
	}
}

//前序遍历
func preorder(node *TreeNode) []*TreeNode {
	var nodes []*TreeNode
	if node != nil {
		nodes = append(nodes, node)
		nodes = append(nodes, preorder(node.Left)...)
		nodes = append(nodes, preorder(node.Right)...)
	}
	return nodes
}

//1、将左子树插入到右子树的地方
//2、将原来的右子树接到左子树的最右边节点【左子树的最右边节点，是右子树的父节点】
//3、考虑新的右子树的根节点，一直重复上边的过程，直到新的右子树为 null
func flattenM2(root *TreeNode) {
	for root != nil {
		//左子树为空是不进行处理
		if root.Left == nil {
			root = root.Right
			continue
		}

		pre := root.Left
		//寻找左子树的最右边节点
		for pre.Right != nil {
			pre = pre.Right
		}
		//将右子树 插在 左子树最右节点 的右节点【步骤2】
		pre.Right = root.Right
		//将左子树，移动到右子树的位置【步骤1】
		root.Right = root.Left
		//根节点的左子树记得清空
		root.Left = nil
		//改变根节点为 右子树根节点
		root = root.Right

	}

}
