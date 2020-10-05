package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=621 lang=golang
 *
 * [621] Task Scheduler
 */

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	if len(tasks) == 0 {
		return 0
	}

	l := len(tasks)
	cntMap := make(map[byte]int)
	for i := 0; i < l; i++ {
		cntMap[tasks[i]]++
	}

	cnts := []int{}
	for _, val := range cntMap {
		cnts = append(cnts, val)
	}

	sort.Slice(cnts, func(i, j int) bool { return cnts[i] > cnts[j] })
	ans := (cnts[0] - 1) * (n + 1)
	for i := 0; i < len(cnts); i++ {
		if cnts[i] == cnts[0] {
			ans++
		} else {
			break
		}
	}

	if len(tasks) > ans {
		ans = len(tasks)
	}
	return ans
}

// @lc code=end

func main() {
	tasks := []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	n := 2
	fmt.Println(leastInterval(tasks, n))
}
