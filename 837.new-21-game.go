package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=837 lang=golang
 *
 * [837] New 21 Game
 */
func new21Game(N int, K int, W int) float64 {
	f := make([]float64, N+1)
	f[0] = 1
	sum := float64(0)
	if K > 0 {
		sum = 1
	}

	for i := 1; i <= N; i++ {
		f[i] = sum / float64(W)
		if i < K {
			sum += f[i]
		}
		if i >= W {
			sum -= f[i-W]
		}
	}
	ans := float64(0)
	for i := K; i <= N; i++ {
		ans += f[i]
	}
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", ans), 64)
	return value
}

// func main() {
// 	fmt.Println(new21Game(21, 17, 10))
// }
