package main

import "fmt"

/*
 * @lc app=leetcode id=386 lang=golang
 *
 * [386] Lexicographical Numbers
 */

// @lc code=start
func check(num, n int, res []int, idx *int) {
	if num > n {
		return
	}

	res[*idx] = num
	*idx++

	num *= 10
	for i := 0; i <= 9; i++ {
		check(num+i, n, res, idx)
	}
}

func lexicalOrder(n int) []int {
	res := make([]int, n)
	idx := 0
	for i := 1; i <= 9; i++ {
		check(i, n, res, &idx)
	}
	return res
}

// @lc code=end

func main() {
	n := 39
	fmt.Println(lexicalOrder(n))
}
