package main

/*
 * @lc app=leetcode id=662 lang=golang
 *
 * [662] Maximum Width of Binary Tree
 */

// Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Depth int
	Num   int
}

func widthOfBinaryTree(root *TreeNode) int {
	ans := 0
	depthMin := make(map[int]int)
	depthMax := make(map[int]int)

	var buildTree func(*TreeNode, int, int) *Node
	buildTree = func(node *TreeNode, depth, num int) *Node {
		if node == nil {
			return nil
		}
		newNode := &Node{node.Val, nil, nil, depth, num}
		if val, ok := depthMin[depth]; ok {
			if num < val {
				depthMin[depth] = num
			}
		} else {
			depthMin[depth] = num
		}

		if val, ok := depthMax[depth]; ok {
			if num > val {
				depthMax[depth] = num
			}
		} else {
			depthMax[depth] = num
		}

		if node.Left != nil {
			newNode.Left = buildTree(node.Left, depth+1, num*2)
		}
		if node.Right != nil {
			newNode.Right = buildTree(node.Right, depth+1, num*2+1)
		}
		return newNode
	}

	buildTree(root, 0, 1)

	for k, v := range depthMin {
		tmp := depthMax[k] - v + 1
		if tmp > ans {
			ans = tmp
		}
	}

	return ans
}

// func main() {
// 	fmt.Println(widthOfBinaryTree(nil))
// }
