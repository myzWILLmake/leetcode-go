package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 */

// @lc code=start
func countPrimes(n int) int {
	primes := map[int]bool{}
	for i := 2; i < n; i++ {
		if !primes[i] {
			for j := 2; j < n/i+1; j++ {
				primes[j*i] = true
			}
		}
	}

	ans := []int{}
	for i := 2; i < n; i++ {
		if !primes[i] {
			ans = append(ans, i)
		}
	}
	return len(ans)
}

// @lc code=end

func main() {
	fmt.Println(countPrimes(6))
}
