package main

import "fmt"

/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	res := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		res[i] = -1
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && res[i-coins[j]] >= 0 {
				if res[i] == -1 || res[i-coins[j]]+1 < res[i] {
					res[i] = res[i-coins[j]] + 1
				}
			}
		}
	}

	return res[amount]
}

// @lc code=end

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}
