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

func TestLetterCombinations(t *testing.T) {
	res := letterCombinations("23")
	fmt.Println(res)
	res = letterCombinations("")
	fmt.Println(res)
	res = letterCombinations("2")
	fmt.Println(res)
}

func TestGenerateParenthesis(t *testing.T) {
	res := generateParenthesis(1)
	fmt.Println(res)
	res = generateParenthesis(2)
	fmt.Println(res)
	res = generateParenthesis(3)
	fmt.Println(res)
	res = generateParenthesis(4)
	fmt.Println(res)
}

func TestNextPermutation(t *testing.T) {
	var arr []int
	arr = []int{1, 2, 3}
	nextPermutation(arr)
	fmt.Println(arr)
	arr = []int{3, 2, 1}
	nextPermutation(arr)
	fmt.Println(arr)
	arr = []int{1, 1, 5}
	nextPermutation(arr)
	fmt.Println(arr)
	arr = []int{1, 2, 3, 4, 6, 5}
	nextPermutation(arr)
	fmt.Println(arr)
}

func TestSearch(t *testing.T) {
	var res int
	res = search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	fmt.Println(res)
	res = search([]int{4, 5, 6, 7, 0, 1, 2}, 3)
	fmt.Println(res)
	res = search([]int{1}, 3)
	fmt.Println(res)
	res = search([]int{5, 1, 3}, 3)
	fmt.Println(res)
}

func TestSearchRange(t *testing.T) {
	res := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Println(res)
	res = searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
	fmt.Println(res)
	res = searchRange([]int{}, 6)
	fmt.Println(res)
	res = searchRange([]int{5, 7, 7, 8, 10}, 8)
	fmt.Println(res)
}
