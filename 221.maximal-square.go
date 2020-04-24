package main

import "fmt"

/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 */

// @lc code=start
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func maximalSquare(matrix [][]byte) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	if m == 0 {
		return 0
	}

	max := 0
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				res[i][j] = 1
				max = 1
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if res[i][j] != 0 {
				res[i][j] = min(res[i-1][j], min(res[i-1][j-1], res[i][j-1])) + 1
				if res[i][j] > max {
					max = res[i][j]
				}
			}
		}
	}
	return max * max
}

// @lc code=end

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	fmt.Println(maximalSquare(matrix))
}
