package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=388 lang=golang
 *
 * [388] Longest Absolute File Path
 */

// @lc code=start
func lengthLongestPath(input string) int {
	input += "\n"
	dirCurr := []string{}
	tabCnt := 0
	lastChPos := -1
	ans := 0
	for i := 0; i < len(input); i++ {
		c := input[i]
		if c == '\n' {
			name := input[lastChPos:i]
			if tabCnt == len(dirCurr) {
				dirCurr = append(dirCurr, name)
			} else {
				dirCurr[tabCnt] = name
			}

			if strings.Contains(name, ".") {
				filename := ""
				for j := 0; j < tabCnt; j++ {
					filename += dirCurr[j] + "/"
				}
				filename += dirCurr[tabCnt]
				if len(filename) > ans {
					ans = len(filename)
				}
			}

			tabCnt = 0
			lastChPos = -1
		} else if c == '\t' {
			tabCnt++
		} else {
			if lastChPos == -1 {
				lastChPos = i
			}
		}
	}
	return ans
}

// @lc code=end

func main() {
	fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"))
}
