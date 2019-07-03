package main

/*
 * @lc app=leetcode id=565 lang=golang
 *
 * [565] Array Nesting
 */
func arrayNesting(nums []int) int {
	res := 0
	n := len(nums)
	hasSeen := make([]bool, n)
	for i := 0; i < n; i++ {
		if !hasSeen[i] {
			cur := 1
			pos := nums[i]
			hasSeen[i] = true
			for !hasSeen[pos] {
				cur++
				pos = nums[pos]
			}
			if cur > res {
				res = cur
			}
		}
	}
	return res
}

// func main() {
// 	nums := []int{5, 4, 0, 3, 1, 6, 2}
// 	fmt.Println(arrayNesting(nums))
// }
