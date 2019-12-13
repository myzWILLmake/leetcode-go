package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=893 lang=golang
 *
 * [893] Groups of Special-Equivalent Strings
 */

// @lc code=start
func getCharMapValue(s string) string {
	evenMap := map[byte]int{}
	oddMap := map[byte]int{}

	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			evenMap[s[i]]++
		} else {
			oddMap[s[i]]++
		}
	}

	res := ""
	for i := 0; i < 26; i++ {
		if v, ok := evenMap['a'+byte(i)]; ok {
			res += string('a'+byte(i)) + strconv.Itoa(v) + "_"
		}
	}

	res += "|"
	for i := 0; i < 26; i++ {
		if v, ok := oddMap['a'+byte(i)]; ok {
			res += string('a'+byte(i)) + strconv.Itoa(v) + "_"
		}
	}

	return res
}

func numSpecialEquivGroups(A []string) int {
	vMap := map[string]int{}
	for i := 0; i < len(A); i++ {
		s := getCharMapValue(A[i])
		vMap[s]++
	}

	return len(vMap)
}

// @lc code=end

func main() {
	A := []string{"abc", "acb", "bac", "bca", "cab", "cba"}
	fmt.Println(numSpecialEquivGroups(A))
}
