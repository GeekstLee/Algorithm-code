package medium

import "testing"

func TestAddTwoNums(t *testing.T) {
	node1 := &ListNode{
		Val:  9,
		Next: nil,
	}
	node2 := &ListNode{
		Val:  9,
		Next: node1,
	}
	node3 := &ListNode{
		Val:  9,
		Next: node2,
	}
	node4 := &ListNode{
		Val:  9,
		Next: nil,
	}
	node5 := &ListNode{
		Val:  9,
		Next: node4,
	}
	node6 := &ListNode{
		Val:  9,
		Next: node5,
	}
	node7 := &ListNode{
		Val:  9,
		Next: node6,
	}
	node8 := &ListNode{
		Val:  9,
		Next: node7,
	}

	addTwoNumbers(node8, node3)
}
