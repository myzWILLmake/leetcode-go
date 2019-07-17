package main

/*
 * @lc app=leetcode id=450 lang=golang
 *
 * [450] Delete Node in a BST
 */

// Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil || root.Right == nil {
			if root.Left != nil {
				root = root.Left
			} else {
				root = root.Right
			}
		} else {
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			root.Val = cur.Val
			root.Right = deleteNode(root.Right, cur.Val)
		}
	}
	return root
}
