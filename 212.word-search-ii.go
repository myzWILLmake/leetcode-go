package main

import "fmt"

/*
 * @lc app=leetcode id=212 lang=golang
 *
 * [212] Word Search II
 */
// @lc code=start
type Node struct {
	end  string
	next []*Node
}

func newNode() *Node {
	node := Node{}
	node.end = ""
	node.next = make([]*Node, 26)
	return &node
}

func check(board [][]byte, x, y int, node *Node, res *[]string) {
	if node.end != "" {
		*res = append(*res, node.end)
		node.end = ""
	}

	n := len(board)
	m := len(board[0])
	if x < 0 || x >= n || y < 0 || y >= m {
		return
	}

	if board[x][y] == '#' {
		return
	}

	idx := board[x][y] - 'a'
	if node.next[idx] == nil {
		return
	}

	node = node.next[idx]
	c := board[x][y]
	board[x][y] = '#'
	check(board, x-1, y, node, res)
	check(board, x, y-1, node, res)
	check(board, x+1, y, node, res)
	check(board, x, y+1, node, res)
	board[x][y] = c
}

func findWords(board [][]byte, words []string) []string {
	n := len(board)
	m := len(board[0])
	k := len(words)
	if n == 0 || m == 0 || k == 0 {
		return []string{}
	}

	root := newNode()
	for _, word := range words {
		node := root
		for _, c := range word {
			idx := c - 'a'
			if node.next[idx] == nil {
				node.next[idx] = newNode()
			}
			node = node.next[idx]
		}
		node.end = word
	}

	output := []string{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			check(board, i, j, root, &output)
		}
	}

	return output
}

// @lc code=end

func main() {
	board := [][]byte{
		{'a', 'b'},
		{'a', 'a'},
	}

	words := []string{"aba", "baa", "bab", "aaab", "aaa", "aaaa", "aaba"}
	fmt.Println(findWords(board, words))
}
