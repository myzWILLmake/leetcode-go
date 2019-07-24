package main

type Node struct {
	id        int
	degree    int
	level     int
	connected []*Node
}

func findMaxNode(id int, nodes []*Node) *Node {
	n := len(nodes)

	for _, node := range nodes {
		node.level = -1
	}

	queue := make([]*Node, n)
	queue[0] = nodes[id]
	nodes[id].level = 0

	head := 0
	tail := 1
	maxLevel := 0
	var maxNode *Node
	for head < tail {
		node := queue[head]
		head++
		for _, nextNode := range node.connected {
			if nextNode.level == -1 {
				nextNode.level = node.level + 1
				if nextNode.level > maxLevel {
					maxLevel = nextNode.level
					maxNode = nextNode
				}
				queue[tail] = nextNode
				tail++
			}
		}
	}
	return maxNode
}

/*
 * @lc app=leetcode id=310 lang=golang
 *
 * [310] Minimum Height Trees
 */
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	if n == 2 {
		return []int{0, 1}
	}

	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &Node{i, 0, -1, []*Node{}}
	}

	for _, edge := range edges {
		idx := edge[0]
		idy := edge[1]
		nodes[idx].degree++
		nodes[idx].connected = append(nodes[idx].connected, nodes[idy])
		nodes[idy].degree++
		nodes[idy].connected = append(nodes[idy].connected, nodes[idx])
	}

	startNode := findMaxNode(0, nodes)
	targetNode := findMaxNode(startNode.id, nodes)
	ans := make(map[int]bool)
	visited := make([]bool, n)
	checkMaxPath(ans, visited, startNode, targetNode, targetNode.level)

	keys := []int{}
	for key := range ans {
		keys = append(keys, key)
	}
	return keys
}

func checkMaxPath(ans map[int]bool, visited []bool, currNode, endNode *Node, maxLevel int) bool {
	visited[currNode.id] = true
	if currNode == endNode {
		return true
	}

	for _, nextNode := range currNode.connected {
		if !visited[nextNode.id] && checkMaxPath(ans, visited, nextNode, endNode, maxLevel) {
			if maxLevel%2 == 0 {
				if currNode.level == maxLevel/2 {
					ans[currNode.id] = true
				}
			} else {
				if currNode.level == maxLevel/2 || currNode.level == maxLevel/2+1 {
					ans[currNode.id] = true
				}
			}
			return true
		}
	}
	return false
}

// func main() {
// 	n := 2
// 	edges := [][]int{{1, 0}}
// 	fmt.Println(findMinHeightTrees(n, edges))
// }
