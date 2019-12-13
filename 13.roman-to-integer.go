package main

import "fmt"

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	vMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	ans := 0
	subHappened := false
	for i := 0; i < len(s); i++ {
		if subHappened {
			subHappened = false
		} else {
			c := s[i]
			if i != len(s)-1 {
				nc := s[i+1]
				if c == 'I' {
					if nc == 'V' {
						ans += 4
						subHappened = true
					} else if nc == 'X' {
						ans += 9
						subHappened = true
					}
				} else if c == 'X' {
					if nc == 'L' {
						ans += 40
						subHappened = true
					} else if nc == 'C' {
						ans += 90
						subHappened = true
					}
				} else if c == 'C' {
					if nc == 'D' {
						ans += 400
						subHappened = true
					} else if nc == 'M' {
						ans += 900
						subHappened = true
					}
				}
			}
			if !subHappened {
				ans += vMap[c]
			}
		}
	}

	return ans
}

// @lc code=end

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}
