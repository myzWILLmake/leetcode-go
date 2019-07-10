package main

/*
 * @lc app=leetcode id=376 lang=golang
 *
 * [376] Wiggle Subsequence
 */
func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	ans := 1
	lastDiff := 0
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff != 0 {
			if lastDiff == 0 {
				ans++
				lastDiff = diff
			} else if (diff > 0) != (lastDiff > 0) {
				ans++
				lastDiff = diff
			}
		}
	}
	return ans
}

// func main() {
// 	nums := []int{1, 1, 1, 1, 2}
// 	fmt.Println(wiggleMaxLength(nums))
// }
