package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode id=492 lang=golang
 *
 * [492] Construct the Rectangle
 */

// @lc code=start
func constructRectangle(area int) []int {
	l := area
	r := 1
	n := int(math.Floor(math.Sqrt(float64(area))))
	for i := 1; i <= n; i++ {
		if area%i == 0 {
			l = area / i
			r = i
		}
	}
	return []int{l, r}
}

// @lc code=end

func main() {
	num := 19
	fmt.Println(constructRectangle(num))
}
