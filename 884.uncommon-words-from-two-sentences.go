package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=884 lang=golang
 *
 * [884] Uncommon Words from Two Sentences
 */
func uncommonFromSentences(A string, B string) []string {
	dictA := make(map[string]int)
	dictB := make(map[string]int)

	wordsA := strings.Split(strings.Trim(A, " "), " ")
	wordsB := strings.Split(strings.Trim(B, " "), " ")

	for _, word := range wordsA {
		dictA[word]++
	}

	for _, word := range wordsB {
		dictB[word]++
	}

	ans := []string{}

	for word, count := range dictA {
		if count == 1 && dictB[word] == 0 {
			ans = append(ans, word)
		}
	}

	for word, count := range dictB {
		if count == 1 && dictA[word] == 0 {
			ans = append(ans, word)
		}
	}

	return ans
}

// func main() {
// 	A := "apple apple"
// 	B := "banana"

// 	fmt.Println(uncommonFromSentences(A, B))
// }
