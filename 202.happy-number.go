package main

import "fmt"

/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func isHappy(n int) bool {
	numMap := make(map[int]bool)
	current := n
	for !numMap[current] {
		numMap[current] = true
		x := current
		current = 0
		digit := 0
		for x/10 > 0 {
			digit = x % 10
			x = x / 10
			current += digit * digit
		}
		current += x * x
		if current == 1 {
			return true
		}
	}
	return false
}

// @lc code=end

func main() {
	n := 19
	fmt.Println(isHappy(n))
}
