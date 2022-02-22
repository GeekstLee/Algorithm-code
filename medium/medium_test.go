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
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr[:len(arr)-1])
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

func TestCombinationSum(t *testing.T) {
	res := combinationSum([]int{2, 3, 5}, 8)
	for _, arr := range res {
		fmt.Println(arr)
	}
	fmt.Println("============")
	res = combinationSum([]int{2, 3, 6, 7}, 7)
	for _, arr := range res {
		fmt.Println(arr)
	}
	fmt.Println("============")
	res = combinationSum([]int{2, 3, 6}, 1)
	for _, arr := range res {
		fmt.Println(arr)
	}
}

func TestPermute(t *testing.T) {
	res := permute([]int{1, 2, 3})
	for _, ans := range res {
		fmt.Println(ans)
	}
	fmt.Println("===========")
	res = permute([]int{1, 0})
	for _, ans := range res {
		fmt.Println(ans)
	}
}

func TestRotate(t *testing.T) {
	var res [][]int
	res = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(res)
	for _, arr := range res {
		fmt.Println(arr)
	}
	fmt.Println("==========")
	res = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(res)
	for _, arr := range res {
		fmt.Println(arr)
	}
}

func TestGroupAnagrams(t *testing.T) {
	res := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	for _, str := range res {
		fmt.Println(str)
	}
	fmt.Println("==========")
	res = groupAnagrams([]string{})
	for _, str := range res {
		fmt.Println(str)
	}
	fmt.Println("==========")
	res = groupAnagrams([]string{"e"})
	for _, str := range res {
		fmt.Println(str)
	}
}

func TestCanJump(t *testing.T) {
	res := canJump([]int{2, 3, 1, 1, 4})
	fmt.Println(res)
	res = canJump([]int{3, 2, 1, 0, 4})
	fmt.Println(res)
}

func TestMerge(t *testing.T) {
	res := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	for _, arr := range res {
		fmt.Println(arr)
	}
	fmt.Println("=======")
	res = merge([][]int{{1, 4}, {4, 5}})
	for _, arr := range res {
		fmt.Println(arr)
	}
}

func TestUniquePaths(t *testing.T) {
	res := uniquePaths(3, 7)
	fmt.Println(res)
	res = uniquePaths(3, 3)
	fmt.Println(res)
}

func TestMinPathSum(t *testing.T) {
	res := minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})
	fmt.Println(res)
	res = minPathSum([][]int{{1, 2, 3}, {4, 5, 6}})
	fmt.Println(res)
}

func TestSortColors(t *testing.T) {
	arr := []int{2, 0, 2, 1, 1, 0}
	sortColors(arr)
	fmt.Println(arr)
}

func TestSubsets(t *testing.T) {
	res := subsets([]int{1, 2, 3, 4})
	for _, arr := range res {
		fmt.Println(arr)
	}
	fmt.Println("===============")
	res = subsets([]int{9, 0, 3, 5, 7})
	for _, arr := range res {
		fmt.Println(arr)
	}
}

func TestExist(t *testing.T) {
	res := exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED")
	fmt.Println(res)
	res = exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE")
	fmt.Println(res)
	res = exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB")
	fmt.Println(res)
}

func TestNumTrees(t *testing.T) {
	res := numTrees(3)
	fmt.Println(res)
	res = numTrees(4)
	fmt.Println(res)
	res = numTrees(5)
	fmt.Println(res)
}

func TestLongestConsecutive(t *testing.T) {
	res := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	fmt.Println(res)
	res = longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	fmt.Println(res)
}

func TestWordBreak(t *testing.T) {
	res := wordBreak("leetcode", []string{"leet", "code"})
	fmt.Println(res)
	res = wordBreak("applepenapple", []string{"apple", "pen"})
	fmt.Println(res)
	res = wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
	fmt.Println(res)
}

func TestMergeList(t *testing.T) {
	l3 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  3,
		Next: l3,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}
	l6 := &ListNode{
		Val:  7,
		Next: nil,
	}
	l5 := &ListNode{
		Val:  4,
		Next: l6,
	}
	l4 := &ListNode{
		Val:  2,
		Next: l5,
	}
	head := mergeList(l1, l4)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func TestMaxProduct(t *testing.T) {
	res := maxProduct([]int{2, 3, -2, 4})
	fmt.Println(res)
	res = maxProduct([]int{-2, 0, -1})
	fmt.Println(res)
}

func TestRob(t *testing.T) {
	res := rob([]int{1, 2, 3, 1})
	fmt.Println(res)
	res = rob([]int{2, 7, 9, 3, 1})
	fmt.Println(res)
}

func TestNumIslands(t *testing.T) {
	res := numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}})
	fmt.Println(res)
	res = numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}})
	fmt.Println(res)
}

func TestCanFinish(t *testing.T) {
	res := canFinish(2, [][]int{{1, 0}})
	fmt.Println(res)
	res = canFinish(2, [][]int{{1, 0}, {0, 1}})
	fmt.Println(res)
	res = canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}})
	fmt.Println(res)
}

func TestFindKthLargest(t *testing.T) {
	res := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	fmt.Println(res)
}

func TestMaximalSquare(t *testing.T) {
	res := maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}})
	fmt.Println(res)
}

func TestLengthOfLIS(t *testing.T) {
	res := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	fmt.Println(res)
	res = lengthOfLIS([]int{0, 1, 0, 3, 2, 3})
	fmt.Println(res)
	res = lengthOfLIS([]int{7, 7, 7, 7, 7})
	fmt.Println(res)
}

func TestProductExceptSelf(t *testing.T) {
	res := productExceptSelf([]int{1, 2, 3, 4})
	fmt.Println(res)
}
