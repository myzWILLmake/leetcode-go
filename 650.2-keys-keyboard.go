package main

import "fmt"

/*
 * @lc app=leetcode id=650 lang=golang
 *
 * [650] 2 Keys Keyboard
 */

// @lc code=start
func minSteps(n int) int {
	if n == 1 {
		return 0
	}

	ans := 0
	d := 2
	for n > 1 {
		for n%d == 0 {
			ans += d
			n /= d
		}
		d++
	}

	return ans
}

// @lc code=end

func main() {
	n := 2
	fmt.Println(minSteps(n))
}
