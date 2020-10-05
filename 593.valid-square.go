package main

import "fmt"

/*
 * @lc app=leetcode id=593 lang=golang
 *
 * [593] Valid Square
 */

// @lc code=start
func dist(p1, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}

func check(p1, p2, p3, p4 []int) bool {
	return dist(p1, p2) == dist(p3, p4) && dist(p1, p4) == dist(p2, p3) &&
		dist(p1, p3) == dist(p2, p4) && dist(p1, p2) == dist(p2, p3) && dist(p1, p2) != 0
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	return check(p1, p2, p3, p4) || check(p1, p3, p2, p4) || check(p1, p3, p4, p2)
}

// @lc code=end

func main() {
	p1 := []int{0, 0}
	p2 := []int{5, 0}
	p3 := []int{5, 4}
	p4 := []int{0, 4}
	fmt.Println(validSquare(p1, p2, p3, p4))
}
