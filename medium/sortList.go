package medium

/**
148. 排序链表
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

 

示例 1：


输入：head = [4,2,1,3]
输出：[1,2,3,4]
示例 2：


输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]
示例 3：

输入：head = []
输出：[]
 

提示：

链表中节点的数目在范围 [0, 5 * 104] 内
-105 <= Node.val <= 105
 

进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func sortList(head *ListNode) *ListNode {
	if head == nil || (head.Next == nil) {
		return head
	}
	//尾节点为nil
	return divideList(head, nil)
}

//链表分隔
func divideList(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	//只有一个元素时
	if head.Next == tail {
		head.Next = nil
		return head
	}
	//快慢指针找链表中点
	slow, fast := head, head
	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}
	//此时slow就是中点
	mid := slow
	//将链表分成两段，注意mid作为边界的情况
	return mergeList(divideList(head, mid), divideList(mid, tail))
}

//升序合并链表
func mergeList(l1, l2 *ListNode) *ListNode {
	var head *ListNode
	if l1.Val <= l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	temp := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}
	if l1 != nil {
		temp.Next = l1
	} else if l2 != nil {
		temp.Next = l2
	}
	return head
}
