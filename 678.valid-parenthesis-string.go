package main

import "fmt"

/*
 * @lc app=leetcode id=678 lang=golang
 *
 * [678] Valid Parenthesis String
 */

// @lc code=start
func checkValidString(s string) bool {
	if s == "" {
		return true
	}
	high := 0
	low := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			high++
			low++
		} else if s[i] == ')' {
			high--
			if low > 0 {
				low--
			}
		} else {
			high++
			if low > 0 {
				low--
			}
		}

		if high < 0 {
			return false
		}
	}

	return low == 0
}

// ((*)
// @lc code=end

func main() {
	s := "(*))"
	fmt.Println(checkValidString(s))
}
