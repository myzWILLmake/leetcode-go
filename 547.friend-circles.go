package main

import "fmt"

/*
 * @lc app=leetcode id=547 lang=golang
 *
 * [547] Friend Circles
 */

// @lc code=start
func getf(x int, f []int) int {
	if f[x] < 0 {
		return x
	}
	f[x] = getf(f[x], f)
	return f[x]
}

func findCircleNum(M [][]int) int {
	n := len(M)
	f := make([]int, n)
	for i := 0; i < n; i++ {
		f[i] = -1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if M[i][j] == 1 {
				fi := getf(i, f)
				fj := getf(j, f)
				if fi != fj {
					if f[fi] <= f[fj] {
						f[fi] += f[fj]
						f[fj] = fi
					} else {
						f[fj] += f[fi]
						f[fi] = fj
					}
				}
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if f[i] < 0 {
			ans++
		}
	}

	return ans
}

// @lc code=end

func main() {
	m := [][]int{{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1}}
	fmt.Println(findCircleNum(m))
}
