package main

import (
	"strconv"
)

/*
 * @lc app=leetcode id=682 lang=golang
 *
 * [682] Baseball Game
 */

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func calPoints(ops []string) int {
	ans := 0
	s := make(stack, 0)
	for _, str := range ops {
		ok := true
		round := 0
		num, err := strconv.Atoi(str)
		if err != nil {
			if str == "+" {
				l := len(s)
				round = s[l-2] + s[l-1]
			} else if str == "C" {
				s, num = s.Pop()
				round = -num
				ok = false
			} else if str == "D" {
				l := len(s)
				round = 2 * s[l-1]
			}
		} else {
			round = num
		}
		if ok {
			s = s.Push(round)
		}
		ans += round
	}

	return ans
}

// func main() {
// 	ops := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
// 	fmt.Println(calPoints(ops))
// }
