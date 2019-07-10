package main

/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 */

// Definition for a binary tree node.
// Treenode ...
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	top := -1

	pushStack := func(s *[]*TreeNode, top *int, node *TreeNode) {
		*top++
		if *top == len(*s) {
			*s = append(*s, node)
		} else {
			(*s)[*top] = node
		}
	}

	popStack := func(s []*TreeNode, top *int) *TreeNode {
		front := s[*top]
		*top--
		return front
	}

	ans := []int{}

	node := root
	for node != nil {
		pushStack(&stack, &top, node)
		node = node.Left
	}

	for top >= 0 {
		n := popStack(stack, &top)
		ans = append(ans, n.Val)

		if n.Right != nil {
			n = n.Right
			for n != nil {
				pushStack(&stack, &top, n)
				n = n.Left
			}
		}
	}

	return ans
}

// func main() {
// 	n3 := TreeNode{3, nil, nil}
// 	n2 := TreeNode{2, &n3, nil}
// 	n1 := TreeNode{1, nil, &n2}
// 	fmt.Println(inorderTraversal(&n1))
// }
