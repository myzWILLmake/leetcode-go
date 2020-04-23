package main

import "fmt"

/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 */

// @lc code=start

func dfs(remain int, limit int, index int, curr []int, res *[][]int) {
	if remain == 0 && limit == 0 {
		dump := make([]int, len(curr))
		copy(dump, curr)
		*res = append(*res, dump)
		return
	}
	if remain < 0 || limit < 0 {
		return
	}

	for i := index; i <= 9; i++ {
		if i > remain {
			break
		}
		curr = append(curr, i)
		dfs(remain-i, limit-1, i+1, curr, res)
		curr = curr[:len(curr)-1]
	}
}

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	curr := []int{}
	dfs(n, k, 1, curr, &res)
	return res
}

// @lc code=end

func main() {
	k := 3
	n := 9
	fmt.Println(combinationSum3(k, n))
}
