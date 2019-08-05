package main

/*
 * @lc app=leetcode id=968 lang=golang
 *
 * [968] Binary Tree Cameras
 */

// Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// Strict Mode 0
// Normal Mode 1
// Camera Mode 2

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func solve(node *TreeNode) (int, int, int) {
	if node == nil {
		return 0, 0, 1001
	}

	l0, l1, l2 := solve(node.Left)
	r0, r1, r2 := solve(node.Right)

	n0 := l1 + r1
	n1 := min(l2+min(r1, r2), r2+min(l1, l2))
	n2 := 1 + min(l0, min(l1, l2)) + min(r0, min(r1, r2))
	return n0, n1, n2
}

func minCameraCover(root *TreeNode) int {
	_, n1, n2 := solve(root)
	return min(n1, n2)
}

// func main() {
// 	node0 := &TreeNode{0, nil, nil}
// 	node1 := &TreeNode{1, node0, nil}
// 	node2 := &TreeNode{2, node1, nil}
// 	node3 := &TreeNode{3, nil, node2}
// 	node4 := &TreeNode{4, nil, node3}
// 	node5 := &TreeNode{5, node4, nil}
// 	node6 := &TreeNode{6, nil, node5}
// 	node7 := &TreeNode{7, nil, nil}
// 	node8 := &TreeNode{8, nil, node7}
// 	node9 := &TreeNode{9, node8, node6}
// 	node10 := &TreeNode{10, nil, node9}
// 	fmt.Println(minCameraCover(node10))
// }
