package main

import "fmt"

/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */

// @lc code=start
func decodeString(s string) string {
	st := []string{}
	tmps := ""
	nt := []int{}
	tmpn := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			st = append(st, tmps)
			tmps = ""
			nt = append(nt, tmpn)
			tmpn = 0
		} else if s[i] == ']' {
			n := nt[len(nt)-1]
			nt = nt[:len(nt)-1]
			ns := ""
			for i := 0; i < n; i++ {
				ns += tmps
			}
			tmps = st[len(st)-1]
			st = st[:len(st)-1]
			tmps += ns
		} else if s[i] >= '0' && s[i] <= '9' {
			tmpn = tmpn*10 + int(s[i]-'0')
		} else {
			tmps += string(s[i])
		}

	}
	return tmps
}

// @lc code=end

func main() {
	input := "2[a3[d2[w]f]2[c]]"
	fmt.Println(decodeString(input))
}
