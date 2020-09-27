package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=274 lang=golang
 *
 * [274] H-Index
 */

// @lc code=start
func hIndex(c []int) int {
	sort.Ints(c)
	n := len(c)
	if n == 0 || c[0] >= n {
		return n
	}
	l := 0
	r := n
	for l+1 < r {
		x := (r - l) / 2
		if c[l+x] >= n-l-x {
			r = l + x
		} else {
			l = l + x
		}
	}

	return n - r
}

// @lc code=end

//
// 12, 3, 0, 2, 5
// 0, 2, 2, 5, 12
// 0, 1, 3, 3, 4, 5, 6

func main() {
	c := []int{}
	fmt.Println(hIndex(c))
}
