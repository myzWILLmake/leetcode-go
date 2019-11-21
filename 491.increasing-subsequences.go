package main

import "fmt"

/*
 * @lc app=leetcode id=491 lang=golang
 *
 * [491] Increasing Subsequences
 */

// @lc code=start
func addAns(hashMap map[string]bool, ans [][]int, seq []int) [][]int {
	id := ""
	for _, num := range seq {
		id += string(num+200) + " "
	}
	if !hashMap[id] {
		ans = append(ans, seq)
		hashMap[id] = true
	}
	return ans
}

func findSubsequences(nums []int) [][]int {
	hashMap := make(map[string]bool)
	ans := make([][]int, 0)
	res := make([][][]int, len(nums))
	for i := 1; i < len(nums); i++ {
		res[i] = make([][]int, 0)
		for j := 0; j < i; j++ {
			if nums[j] <= nums[i] {
				res[i] = append(res[i], []int{nums[j], nums[i]})
				ans = addAns(hashMap, ans, []int{nums[j], nums[i]})
				for _, seq := range res[j] {
					newSeq := make([]int, len(seq))
					copy(newSeq, seq)
					newSeq = append(newSeq, nums[i])
					res[i] = append(res[i], newSeq)
					ans = addAns(hashMap, ans, newSeq)
				}
			}
		}
	}
	return ans
}

// @lc code=end

func main() {
	nums := []int{1, 2, 3, 4, 5, 5}
	fmt.Println(findSubsequences(nums))
}
