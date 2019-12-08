package main

import "fmt"

/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func inorderTraversal(node *TreeNode, nums []int) []int {
	if node.Left != nil {
		nums = inorderTraversal(node.Left, nums)
	}
	nums = append(nums, node.Val)
	if node.Right != nil {
		nums = inorderTraversal(node.Right, nums)
	}
	return nums
}

func getMinimumDifference(root *TreeNode) int {
	nums := []int{}
	nums = inorderTraversal(root, nums)
	ans := int(^uint(0) >> 1)
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff < ans {
			ans = diff
		}
	}
	return ans
}

// @lc code=end

func main() {
	// node2 := &TreeNode{40, nil, nil}
	node3 := &TreeNode{100, nil, nil}
	node1 := &TreeNode{1, nil, node3}
	fmt.Println(getMinimumDifference(node1))
}
