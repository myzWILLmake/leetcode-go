package main

/*
 * @lc app=leetcode id=673 lang=golang
 *
 * [673] Number of Longest Increasing Subsequence
 */

// @lc code=start
func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxLength := 1
	length := make([]int, len(nums))
	count := make([]int, len(nums))
	length[0] = 1
	count[0] = 1
	for i := 1; i < len(nums); i++ {
		max := 0
		cnt := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if length[j] > max {
					max = length[j]
					cnt = count[j]
				} else if length[j] == max {
					cnt += count[j]
				}
			}
		}
		length[i] = max + 1
		if cnt == 0 {
			count[i] = 1
		} else {
			count[i] = cnt
		}
		if length[i] > maxLength {
			maxLength = length[i]
		}
	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		if length[i] == maxLength {
			ans += count[i]
		}
	}
	return ans
}

// @lc code=end

// func main() {
// 	nums := []int{1}
// 	fmt.Println(findNumberOfLIS(nums))
// }
