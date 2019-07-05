package main

/*
 * @lc app=leetcode id=126 lang=golang
 *
 * [126] Word Ladder II
 */
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	checkEdge := func(a, b string) bool {
		diff := 0
		for idx := range a {
			if a[idx] != b[idx] {
				diff++
			}
		}
		if diff != 1 {
			return false
		}
		return true
	}

	type Node struct {
		idx   int
		word  string
		next  []*Node
		paths [][]string
	}

	ans := [][]string{}

	n := len(wordList)
	nodes := make([]*Node, n)
	endWordNo := n
	for i := 0; i < n; i++ {
		node := Node{i, wordList[i], []*Node{}, [][]string{}}
		nodes[i] = &node
		if endWord == wordList[i] {
			endWordNo = i
		}
	}
	if endWordNo == n {
		return ans
	}

	// root as No.n node
	root := &Node{n, beginWord, []*Node{}, [][]string{{}}}
	for i := 0; i < n; i++ {
		nx := nodes[i]
		for j := 0; j < i; j++ {
			ny := nodes[j]
			if checkEdge(nx.word, ny.word) {
				nx.next = append(nx.next, ny)
				ny.next = append(ny.next, nx)
			}
		}
		if checkEdge(root.word, nx.word) {
			root.next = append(root.next, nx)
		}
	}

	if len(root.next) == 0 {
		return ans
	}

	// 0 not visited, 1 added but not processed, 2 processed
	visited := make([]int, n+1)
	queue := []*Node{root}
	head := 0
	ok := false
	for head < len(queue) {
		top := queue[head]
		head++
		visited[top.idx] = 2
		for idx, path := range top.paths {
			top.paths[idx] = append(path, top.word)
		}
		if top.idx == endWordNo {
			ok = true
			break
		}
		for _, node := range top.next {
			if visited[node.idx] == 0 {
				queue = append(queue, node)
				visited[node.idx] = 1
			}
			if visited[node.idx] != 2 {
				if len(node.paths) == 0 || len(top.paths[0]) == len(node.paths[0]) {
					for _, path := range top.paths {
						newPath := make([]string, len(path))
						copy(newPath, path)
						node.paths = append(node.paths, newPath)
					}
				}
			}
		}
	}

	if ok {
		ans = nodes[endWordNo].paths
	}

	return ans
}

// func main() {
// 	beginWord := "hit"
// 	endWord := "cog"
// 	wordList := []string{"hot", "dot", "dog", "lot", "log"}

// 	fmt.Println(findLadders(beginWord, endWord, wordList))
// }
