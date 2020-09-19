package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode id=400 lang=golang
 *
 * [400] Nth Digit
 */

// @lc code=start

func pick(num, nth, base int) int {
	base = base - nth - 1
	x := int(math.Pow10(base))
	return (num / x) % 10
}

func findNthDigit(n int) int {
	base := 1
	mul := 9

	if n < 10 {
		return n
	}

	for n > base*mul {
		n -= base * mul
		base++
		mul *= 10
	}

	n -= 1
	num := n / base
	rmd := n % base

	num += int(math.Pow10(base - 1))
	return pick(num, rmd, base)
}

// @lc code=end

func main() {
	n := 2889
	fmt.Println(findNthDigit(n))
}

// 1 * 9
// 2 * 90
// 3 * 900
// 4 * 9000
// 5 * 90000
