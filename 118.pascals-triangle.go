package main

import "fmt"

/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

// @lc code=start
func generate(numRows int) [][]int {
	ans := make([][]int, 0)

	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		if i > 0 {
			lastRow := ans[i-1]
			for j := 1; j < i; j++ {
				row[j] = lastRow[j-1] + lastRow[j]
			}
		}
		ans = append(ans, row)
	}

	return ans
}

// @lc code=end

func main() {
	fmt.Println(generate(10))
}
