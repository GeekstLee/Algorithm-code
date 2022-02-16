package medium

import (
	"fmt"
	"testing"
)

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
func TestLengthOfLongestSubstring(t *testing.T) {
	var res int
	res = lengthOfLongestSubstring("abcabcbb")
	fmt.Println("abcabcbb==", res)
	res = lengthOfLongestSubstring("aaa")
	fmt.Println("aaa==", res)
	res = lengthOfLongestSubstring("")
	fmt.Println("==", res)
	res = lengthOfLongestSubstring("abcba")
	fmt.Println("abcba==", res)
	res = lengthOfLongestSubstring("pwwkew")
	fmt.Println("pwwkew==", res)
	res = lengthOfLongestSubstring("abcdefg")
	fmt.Println("abcdefg==", res)
}
func TestMy(t *testing.T) {
	str := "2sdad"
	fmt.Println(str[0:1])
}

func TestLongestPalindrome(t *testing.T) {
	res := longestPalindrome("babad")
	fmt.Println(res)
	res = longestPalindrome("cbbd")
	fmt.Println(res)

	res = longestPalindromeM2("babad")
	fmt.Println(res)
	res = longestPalindromeM2("cbbd")
	fmt.Println(res)
}
func TestMaxArea(t *testing.T) {
	var res int
	res = maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	println(res)
	res = maxArea([]int{1, 1})
	println(res)
}

func TestThreeSum(t *testing.T) {
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	for _, nums := range res {
		fmt.Println(nums)
	}
	fmt.Println("========")
	res = threeSum([]int{0, 0, 0, 0, 0})
	for _, nums := range res {
		fmt.Println(nums)
	}
}
