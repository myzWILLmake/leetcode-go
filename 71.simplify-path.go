package main

import "fmt"

/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start
func simplifyPath(path string) string {
	path += "/"
	dirStack := []string{}
	var lastChar byte
	lastChar = '/'
	name := ""
	for i := 0; i < len(path); i++ {
		ch := path[i]
		if ch == '/' {
			if lastChar != '/' {
				if name == ".." {
					l := len(dirStack)
					if l > 0 {
						dirStack = dirStack[0 : l-1]
					}
				} else if name != "." {
					dirStack = append(dirStack, name)
				}
				name = ""
			}
		} else {
			name += string(ch)
		}
		lastChar = ch
	}

	if len(dirStack) == 0 {
		return "/"
	}

	ans := ""
	for i := 0; i < len(dirStack); i++ {
		ans += "/" + dirStack[i]
	}
	return ans
}

// @lc code=end

func main() {
	path := "/a//b////c/d//././/.."
	fmt.Println(simplifyPath(path))
}
