package medium

/**
19. 删除链表的倒数第 N 个结点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

 

示例 1：


输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]
 

提示：

链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	//先遍历得到链表长度
	Len := 0
	temp := head
	for {
		if temp != nil {
			Len++
		} else {
			break
		}
		temp = temp.Next
	}
	pre := &ListNode{
		Val:  0,
		Next: head,
	}
	deleteNum := Len - n //要删除第几个节点
	temp = pre
	//pre->1->2->3->4
	for i := 0; i < deleteNum; i++ {
		temp = temp.Next
	}
	temp.Next = temp.Next.Next
	return pre.Next

}

//可以使用两个指针 first 和 second 同时对链表进行遍历
//first 比 second 超前 n 个节点
//当 first 遍历到链表的末尾时，second 就恰好处于倒数第 n 个节点
func removeNthFromEndM2(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	pre := &ListNode{
		Val:  0,
		Next: head,
	}
	dummy := pre
	first := head
	for i := 0; i < n; i++ {
		first = first.Next
	}

	for {
		if first != nil {
			first = first.Next
			pre = pre.Next
		} else {
			break
		}
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}
