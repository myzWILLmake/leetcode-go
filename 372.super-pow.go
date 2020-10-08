package main

import "fmt"

/*
 * @lc app=leetcode id=372 lang=golang
 *
 * [372] Super Pow
 */

// @lc code=start

func spow(a int, b int) int {
	if b == 0 {
		return 1
	}
	a0 := a % 1337
	a = a0
	for i := 1; i < b; i++ {
		a = (a * a0) % 1337
	}
	return a
}

func superPow(a int, b []int) int {
	if len(b) == 1 && b[0] == 0 {
		return 1
	}
	n := len(b)
	rb := make([]int, n)
	for i := 0; i < n; i++ {
		rb[i] = b[len(b)-i-1]
	}
	base := a % 1337
	res := 1
	for i := 0; i < n; i++ {
		cur := spow(base, rb[i])
		res = (cur * res) % 1337
		base = spow(base, 10)
	}
	return res
}

// @lc code=end

func main() {
	a := 2147483647
	b := []int{2, 0, 0}
	fmt.Println(superPow(a, b))
}
