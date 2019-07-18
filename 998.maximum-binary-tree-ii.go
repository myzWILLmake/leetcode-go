package main

/*
 * @lc app=leetcode id=998 lang=golang
 *
 * [998] Maximum Binary Tree II
 */

// Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxIdx := -1
	max := 0
	for idx, num := range nums {
		if maxIdx == -1 || num > max {
			maxIdx = idx
			max = num
		}
	}

	newNode := TreeNode{max, nil, nil}
	if maxIdx > 0 {
		newNode.Left = constructMaximumBinaryTree(nums[:maxIdx])
	}

	if maxIdx < len(nums)-1 {
		newNode.Right = constructMaximumBinaryTree(nums[maxIdx+1:])
	}

	return &newNode
}

func getNumsFromTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var left, right []int
	left = getNumsFromTree(root.Left)
	right = getNumsFromTree(root.Right)

	nums := make([]int, len(left)+len(right)+1)
	copy(nums, left)
	nums[len(left)] = root.Val
	copy(nums[len(left)+1:], right)

	return nums
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	nums := getNumsFromTree(root)
	newNums := make([]int, len(nums)+1)
	copy(newNums, nums)
	newNums[len(nums)] = val
	return constructMaximumBinaryTree(newNums)
}

// func printTree(node *TreeNode) {
// 	if node == nil {
// 		return
// 	}
// 	fmt.Println(node)
// 	printTree(node.Left)
// 	printTree(node.Right)
// }

// func main() {
// 	nums := []int{3, 2, 1, 6, 0, 5}
// 	printTree(insertIntoMaxTree(constructMaximumBinaryTree(nums), 10))
// }
