package main

/*
 * @lc app=leetcode id=701 lang=golang
 *
 * [701] Insert into a Binary Search Tree
 */

// // TreeNode Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := root
	for node != nil {
		if val < node.Val {
			if node.Left == nil {
				newNode := TreeNode{val, nil, nil}
				node.Left = &newNode
				break
			} else {
				node = node.Left
			}
		} else {
			if node.Right == nil {
				newNode := TreeNode{val, nil, nil}
				node.Right = &newNode
				break
			} else {
				node = node.Right
			}
		}
	}
	return root
}
