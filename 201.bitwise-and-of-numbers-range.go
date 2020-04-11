package main

import "fmt"

/*
 * @lc app=leetcode id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 */

// @lc code=start
func rangeBitwiseAnd(m int, n int) int {
	ans := m
	for i := m + 1; i <= n; i++ {
		ans = ans & i
	}
	return ans
}

// @lc code=end

func main() {
	m := 5
	n := 7
	fmt.Println(rangeBitwiseAnd(m, n))
}
