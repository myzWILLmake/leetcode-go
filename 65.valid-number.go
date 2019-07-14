package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=65 lang=golang
 *
 * [65] Valid Number
 */
func isNumber(s string) bool {
	s = strings.Trim(s, " ")

	eIdx := -1
	signIdx := -1
	pointIdx := -1
	firstNumIdx := -1

	for idx, c := range s {
		if c >= '0' && c <= '9' {
			if firstNumIdx == -1 {
				firstNumIdx = idx
			}
		} else if c == 'e' {
			if eIdx != -1 {
				return false
			}
			if firstNumIdx == -1 {
				return false
			}
			eIdx = idx
			pointIdx = -1
			signIdx = -1
			firstNumIdx = -1
		} else if c == '.' {
			if pointIdx != -1 {
				return false
			}
			if eIdx != -1 {
				return false
			}
			pointIdx = idx
		} else if c == '+' || c == '-' {
			if pointIdx != -1 {
				return false
			}
			if firstNumIdx != -1 {
				return false
			}
			if signIdx != -1 {
				return false
			}
			signIdx = idx
		} else {
			return false
		}
	}

	if eIdx == len(s)-1 || signIdx == len(s)-1 {
		return false
	}

	if pointIdx == len(s)-1 && firstNumIdx == -1 {
		return false
	}

	return true
}

// func main() {
// 	fmt.Println(isNumber("32.e-80123"))
// }
