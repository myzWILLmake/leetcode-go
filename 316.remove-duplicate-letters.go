package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=316 lang=golang
 *
 * [316] Remove Duplicate Letters
 */
func removeDuplicateLetters(s string) string {
	cnt := make(map[rune]int)

	for _, c := range s {
		cnt[c]++
	}

	pos := 0
	for idx, c := range s {
		if c < rune(s[pos]) {
			pos = idx
		}
		cnt[c]--
		if cnt[c] == 0 {
			break
		}
	}

	if len(s) == 0 {
		return ""
	}

	next := ""
	if pos+1 != len(s) {
		next = removeDuplicateLetters(strings.Replace(s[pos+1:], string(s[pos]), "", -1))
	}
	return string(s[pos]) + next
}

// func main() {
// 	fmt.Println(removeDuplicateLetters(""))
// }
