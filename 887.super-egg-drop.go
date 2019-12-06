package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=887 lang=golang
 *
 * [887] Super Egg Drop
 */

// @lc code=start
func superEggDrop(K int, N int) int {
	if K == 1 || N == 1 {
		return N
	}
	// f(t, k)= 1 + f(t-1, k-1) + f(t-1, k)
	f := make([][]int, 2)
	f[0] = make([]int, N)
	f[1] = make([]int, N)
	for i := 0; i < N; i++ {
		f[1][i] = i + 1
	}

	ans := N
	for depth := 2; depth <= K; depth++ {
		f = append(f, make([]int, N))
		f[depth][0] = 1
		for i := 1; i < N; i++ {
			tmp := 1 + f[depth][i-1] + f[depth-1][i-1]
			if tmp >= N && i+1 < ans {
				ans = i + 1
				f[depth][i] = tmp
				break
			}
			f[depth][i] = tmp
		}
	}
	return ans
}

// @lc code=end

func main() {
	K := 2
	N := 8129
	fmt.Println(superEggDrop(K, N))
}
