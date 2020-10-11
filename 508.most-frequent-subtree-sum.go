package main

import "fmt"

/*
 * @lc app=leetcode id=508 lang=golang
 *
 * [508] Most Frequent Subtree Sum
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getSum(node *TreeNode, sums map[int]int) int {
	if node == nil {
		return 0
	}
	l := getSum(node.Left, sums)
	r := getSum(node.Right, sums)
	sum := l + r + node.Val
	sums[sum]++
	return sum
}

func findFrequentTreeSum(root *TreeNode) []int {
	sums := make(map[int]int)
	getSum(root, sums)
	max := 0
	ans := []int{}
	for key, times := range sums {
		if times > max {
			ans = []int{}
			max = times
		}
		if times == max {
			ans = append(ans, key)
		}
	}
	return ans
}

// @lc code=end
func main() {
	t1 := &TreeNode{2, nil, nil}
	t2 := &TreeNode{-3, nil, nil}
	t3 := &TreeNode{5, t1, t2}
	fmt.Println(findFrequentTreeSum(t3))
}
