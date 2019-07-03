package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=165 lang=golang
 *
 * [165] Compare Version Numbers
 */
func compareVersion(version1 string, version2 string) int {
	var dealNum = func(s string) int {
		r := 0
		for _, w := range s {
			r = r*10 + int(w-'0')
		}
		return r
	}

	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	n1 := make([]int, len(v1))
	n2 := make([]int, len(v2))

	for i := 0; i < len(v1); i++ {
		n1[i] = dealNum(v1[i])
	}

	for i := 0; i < len(v2); i++ {
		n2[i] = dealNum(v2[i])
	}

	i, j := 0, 0
	for ; i < len(n1) && j < len(n2); i, j = i+1, j+1 {
		if n1[i] > n2[j] {
			return 1
		}
		if n1[i] < n2[j] {
			return -1
		}
	}

	pos := i
	n := n1
	l := len(n1)
	res := 1
	if j < len(n2) {
		pos = j
		n = n2
		l = len(n2)
		res = -1
	}

	for ; pos < l; pos++ {
		if n[pos] != 0 {
			return res
		}
	}
	return 0
}

// func main() {
// 	fmt.Println(compareVersion("1.0.1", "1"))
// }
