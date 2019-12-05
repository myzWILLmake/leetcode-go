package main

import "fmt"

/*
 * @lc app=leetcode id=477 lang=golang
 *
 * [477] Total Hamming Distance
 */

// @lc code=start
func totalHammingDistance(nums []int) int {
	bit1Cnt := make([]int, 32)
	bit0Cnt := make([]int, 32)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for i := 0; i < 32; i++ {
			if num%2 == 0 {
				bit0Cnt[i]++
			} else {
				bit1Cnt[i]++
			}
			num >>= 1
		}
	}

	ans := 0
	for i := 0; i < 32; i++ {
		ans += bit1Cnt[i] * bit0Cnt[i]
	}

	return ans
}

// @lc code=end

func main() {
	nums := []int{4, 14, 2}
	fmt.Println(totalHammingDistance(nums))
}
