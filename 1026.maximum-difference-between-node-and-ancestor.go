package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findMinAndMaxChildrenWithRes(node *TreeNode) (int, int, int) {
	min := node.Val
	max := node.Val
	res := 0

	if node.Left != nil {
		lmin, lmax, lres := findMinAndMaxChildrenWithRes(node.Left)
		if lmin < min {
			min = lmin
		}
		if lmax > max {
			max = lmax
		}
		if lres > res {
			res = lres
		}
	}

	if node.Right != nil {
		rmin, rmax, rres := findMinAndMaxChildrenWithRes(node.Right)
		if rmin < min {
			min = rmin
		}
		if rmax > max {
			max = rmax
		}
		if rres > res {
			res = rres
		}
	}

	if abs(min-node.Val) > res {
		res = abs(min - node.Val)
	}
	if abs(max-node.Val) > res {
		res = abs(max - node.Val)
	}

	return min, max, res
}

func maxAncestorDiff(root *TreeNode) int {
	_, _, res := findMinAndMaxChildrenWithRes(root)
	return res
}

// func main() {
// 	n2 := TreeNode{3, nil, nil}
// 	n3 := TreeNode{10, nil, nil}
// 	n1 := TreeNode{8, &n2, &n3}
// 	fmt.Println(maxAncestorDiff(&n1))
// }
