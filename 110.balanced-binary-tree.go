package main

/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TreeNode ...
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	ans := true
	var getDepth func(*TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node.Left == nil && node.Right == nil {
			return 0
		}
		l := 0
		r := 0
		if node.Left != nil {
			l = getDepth(node.Left) + 1
		}
		if node.Right != nil {
			r = getDepth(node.Right) + 1
		}

		if l > r {
			if l-r > 1 {
				ans = false
			}
			return l
		}

		if r-l > 1 {
			ans = false
		}
		return r
	}
	getDepth(root)
	return ans
}

// func main() {
// 	nw := TreeNode{0, nil, nil}
// 	n15 := TreeNode{15, &nw, nil}
// 	n7 := TreeNode{7, nil, nil}
// 	n20 := TreeNode{20, &n15, &n7}
// 	n9 := TreeNode{9, nil, nil}
// 	n3 := TreeNode{7, &n9, &n20}
// 	fmt.Println(isBalanced(&n3))

// }
