package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=395 lang=golang
 *
 * [395] Longest Substring with At Least K Repeating Characters
 */

// @lc code=start
func longestSubstring(s string, k int) int {
	freq := make(map[byte]int)
	n := len(s)
	for i := 0; i < n; i++ {
		freq[s[i]]++
	}

	leastCnt := k + 1
	var leastByt byte
	for k, v := range freq {
		if v < leastCnt {
			leastCnt = v
			leastByt = k
		}
	}

	if leastCnt < k {
		ss := strings.Split(s, string(leastByt))
		max := 0
		for _, x := range ss {
			res := longestSubstring(x, k)
			if res > max {
				max = res
			}
		}
		return max
	} else {
		return n
	}

}

// @lc code=end

func main() {
	s := "aaabb"
	k := 3
	fmt.Println(longestSubstring(s, k))
}
