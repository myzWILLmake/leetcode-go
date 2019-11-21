package main

import "fmt"

/*
 * @lc app=leetcode id=953 lang=golang
 *
 * [953] Verifying an Alien Dictionary
 */

// @lc code=start

func compareWord(w1, w2 string, order map[byte]int) bool {
	i := 0
	for i < len(w1) && i < len(w2) {
		if order[w1[i]] < order[w2[i]] {
			return true
		} else if order[w1[i]] > order[w2[i]] {
			return false
		} else {
			i++
		}
	}
	if len(w1) <= len(w2) {
		return true
	}
	return false
}

func isAlienSorted(words []string, order string) bool {
	orderMap := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		orderMap[order[i]] = i
	}

	for i := 1; i < len(words); i++ {
		if !compareWord(words[i-1], words[i], orderMap) {
			return false
		}
	}

	return true
}

// @lc code=end

func main() {
	words := []string{"hello", "leetcode"}
	order := "hlabcdefgijkmnopqrstuvwxyz"
	fmt.Println(isAlienSorted(words, order))
}
