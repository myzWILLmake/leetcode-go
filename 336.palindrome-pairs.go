package main

import "fmt"

/*
 * @lc app=leetcode id=336 lang=golang
 *
 * [336] Palindrome Pairs
 */

// @lc code=start

func checkPalindrome(word string) bool {
	i := 0
	j := len(word) - 1
	for i < j {
		if word[i] == word[j] {
			i++
			j--
		} else {
			return false
		}
	}

	return true
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func palindromePairs(words []string) [][]int {
	m := map[string]int{}
	for idx, word := range words {
		m[word] = idx
	}

	ans := [][]int{}
	for idx, word := range words {
		if word == "" {
			for i, w := range words {
				if w != "" && checkPalindrome(w) {
					ans = append(ans, []int{i, idx})
				}
			}
			continue
		}

		rWord := reverseString(word)
		for pos := range rWord {
			pre := rWord[:pos]
			suf := rWord[pos:]
			if idxPre, ok := m[pre]; ok && checkPalindrome(suf) && idx != idxPre {
				ans = append(ans, []int{idxPre, idx})
			}
			if sufPre, ok := m[suf]; ok && checkPalindrome(pre) && sufPre != idx {
				ans = append(ans, []int{idx, sufPre})
			}
		}
	}
	return ans
}

// @lc code=end

func main() {
	words := []string{"abcd", "dcba", "lls", "s", "sssll"}
	fmt.Println(palindromePairs(words))
}
