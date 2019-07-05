package main

/*
 * @lc app=leetcode id=515 lang=golang
 *
 * [515] Find Largest Value in Each Tree Row
 */

// Definition for a binary tree node.
// TreeNode ...
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{}
	queue := []*TreeNode{}
	lvQueue := []int{}

	queue = append(queue, root)
	lvQueue = append(lvQueue, 0)
	head := 0
	for head < len(queue) {
		node := queue[head]
		level := lvQueue[head]
		if level >= len(ans) {
			ans = append(ans, node.Val)
		} else {
			if node.Val > ans[level] {
				ans[level] = node.Val
			}
		}
		head++

		if node.Left != nil {
			queue = append(queue, node.Left)
			lvQueue = append(lvQueue, level+1)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
			lvQueue = append(lvQueue, level+1)
		}

	}
	return ans
}

// func main() {
// 	n1 := TreeNode{5, nil, nil}
// 	n2 := TreeNode{3, nil, nil}
// 	n3 := TreeNode{9, nil, nil}
// 	n4 := TreeNode{3, &n1, &n2}
// 	n5 := TreeNode{2, nil, &n3}
// 	n6 := TreeNode{1, &n4, &n5}
// 	fmt.Println(largestValues(&n6))
// }
