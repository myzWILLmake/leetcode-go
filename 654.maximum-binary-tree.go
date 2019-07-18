package main

/*
 * @lc app=leetcode id=654 lang=golang
 *
 * [654] Maximum Binary Tree
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

// func printTree(node *TreeNode) {
// 	if node == nil {
// 		return
// 	}
// 	fmt.Println(node)
// 	printTree(node.Left)
// 	printTree(node.Right)
// }

// func main() {
// 	nums := []int{3, 2, 1, 6, 0, 5, 7}
// 	printTree(constructMaximumBinaryTree(nums))
// }
