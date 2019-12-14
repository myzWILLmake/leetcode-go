package main

import "fmt"

/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func trap(height []int) int {
	n := len(height)
	lmax := make([]int, n)
	rmax := make([]int, n)

	lm := 0
	for i := 0; i < n; i++ {
		if height[i] > lm {
			lm = height[i]
			lmax[i] = height[i]
		} else {
			lmax[i] = lm
		}
	}

	rm := 0
	for i := n - 1; i >= 0; i-- {
		if height[i] > rm {
			rm = height[i]
			rmax[i] = height[i]
		} else {
			rmax[i] = rm
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res += min(lmax[i], rmax[i]) - height[i]
	}

	return res
}

// @lc code=end

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
