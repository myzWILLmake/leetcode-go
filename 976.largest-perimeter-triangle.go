package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=976 lang=golang
 *
 * [976] Largest Perimeter Triangle
 */
func largestPerimeter(A []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(A)))
	for i := 0; i < len(A)-2; i++ {
		if A[i+1]+A[i+2] > A[i] {
			return A[i+1] + A[i+2] + A[i]
		}
	}
	return 0
}

// func main() {
// 	A := []int{3, 2, 3, 4}
// 	fmt.Println(largestPerimeter(A))
// }
