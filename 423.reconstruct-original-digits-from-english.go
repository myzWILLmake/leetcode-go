package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=423 lang=golang
 *
 * [423] Reconstruct Original Digits from English
 */

// @lc code=start
// zero   z->0
// one    after 0,2,4, o->1
// two	  w->2
// three after 8, h->3
// four   u->4
// five   after 4, f->5
// six   x->6
// seven after 5, v->7
// eight g->8
// nine last->9
func checkDigit(digitCnt []int, charMap map[byte]int, num int, ns string, disChar byte) {
	if charMap[disChar] > 0 {
		cnt := charMap[disChar]
		digitCnt[num] += cnt
		for i := 0; i < len(ns); i++ {
			charMap[ns[i]] -= cnt
		}
	}
}

func originalDigits(s string) string {
	charMap := make(map[byte]int)
	digitCnt := make([]int, 10)
	for i := 0; i < len(s); i++ {
		charMap[s[i]]++
	}

	checkDigit(digitCnt, charMap, 0, "zero", 'z')
	checkDigit(digitCnt, charMap, 2, "two", 'w')
	checkDigit(digitCnt, charMap, 4, "four", 'u')
	checkDigit(digitCnt, charMap, 6, "six", 'x')
	checkDigit(digitCnt, charMap, 8, "eight", 'g')
	checkDigit(digitCnt, charMap, 3, "three", 'h')
	checkDigit(digitCnt, charMap, 5, "five", 'f')
	checkDigit(digitCnt, charMap, 7, "seven", 'v')
	checkDigit(digitCnt, charMap, 1, "one", 'o')
	checkDigit(digitCnt, charMap, 9, "nine", 'i')

	ans := ""
	for i := 0; i < 10; i++ {
		if digitCnt[i] > 0 {
			x := strconv.Itoa(i)
			for j := 0; j < digitCnt[i]; j++ {
				ans += x
			}
		}
	}
	return ans
}

// @lc code=end

func main() {
	s := "fviefuro"
	fmt.Println(originalDigits(s))
}
