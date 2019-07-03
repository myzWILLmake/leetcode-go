package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=912 lang=golang
 *
 * [912] Sort an Array
 */
func sortArray(nums []int) []int {
	sort.Ints(nums)
	return nums
}

// func main() {
// 	nums := []int{5, 2, 3, 1}
// 	fmt.Println(sortArray(nums))
// }
