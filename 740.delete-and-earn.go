package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=740 lang=golang
 *
 * [740] Delete and Earn
 */
func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}

	numKeys := make([]int, len(numMap))
	for k := range numMap {
		numKeys = append(numKeys, k)
	}

	sort.Ints(numKeys)
	f := make([][]int, len(numKeys))
	f[0] = make([]int, 2)
	f[0][0] = 0
	f[0][1] = numMap[numKeys[0]] * numKeys[0]
	ans := max(f[0][0], f[0][1])
	for i := 1; i < len(numKeys); i++ {
		f[i] = make([]int, 2)
		f[i][0] = max(f[i-1][0], f[i-1][1])
		val := numKeys[i] * numMap[numKeys[i]]
		if numKeys[i-1]+1 == numKeys[i] {
			f[i][1] = f[i-1][0] + val
		} else {
			f[i][1] = max(f[i-1][0]+val, f[i-1][1]+val)
		}
		ans = max(f[i][0], f[i][1])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// func main() {
// 	nums := []int{}
// 	fmt.Println(deleteAndEarn(nums))
// }
