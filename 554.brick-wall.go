package main

import "fmt"

/*
 * @lc app=leetcode id=554 lang=golang
 *
 * [554] Brick Wall
 */

// @lc code=start
func leastBricks(wall [][]int) int {
	res := make(map[int]int)
	max := 0
	n := len(wall)
	for i := 0; i < n; i++ {
		b := wall[i]
		sum := 0
		for j := 0; j < len(b)-1; j++ {
			sum += b[j]
			res[sum]++
			if res[sum] > max {
				max = res[sum]
			}
		}
	}
	return n - max
}

// @lc code=end

func main() {
	wall := [][]int{{1, 1}, {2}, {2}, {1, 1}}
	fmt.Println(leastBricks(wall))
}
