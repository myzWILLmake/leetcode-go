package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=791 lang=golang
 *
 * [791] Custom Sort String
 */

// @lc code=start
func customSortString(S string, T string) string {
	order := map[byte]int{}
	for i := 0; i < len(S); i++ {
		order[S[i]] = i
	}

	t := []byte(T)
	sort.Slice(t, func(i, j int) bool { return order[t[i]] < order[t[j]] })
	return string(t)
}

// @lc code=end

func main() {
	s := "cba"
	t := "abcd"
	fmt.Println(customSortString(s, t))
}
