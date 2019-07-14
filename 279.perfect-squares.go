package main

import (
	"math"
)

/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */
func numSquares(n int) int {
	//f[n] = min(f[i]+1) f[n]-f[i] == x^2

	f := make([]int, n+1)
	f[1] = 1
	for i := 2; i <= n; i++ {
		min := f[i-1]
		maxPerfectNum := int(math.Floor(math.Sqrt(float64(i))))
		for j := 2; j <= maxPerfectNum; j++ {
			if f[i-j*j] < min {
				min = f[i-j*j]
			}
		}
		f[i] = min + 1
	}

	return f[n]
}

// func main() {
// 	fmt.Println(numSquares(13))
// }
