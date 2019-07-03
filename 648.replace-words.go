package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=648 lang=golang
 *
 * [648] Replace Words
 */
func replaceWords(dict []string, sentence string) string {
	type TrieNode struct {
		word     string
		next     []*TrieNode
		finished bool
	}

	var makeNode = func() *TrieNode {
		node := TrieNode{}
		node.next = make([]*TrieNode, 26)
		node.finished = false
		return &node
	}

	root := makeNode()
	for _, s := range dict {
		node := root
		word := ""
		for _, w := range s {
			if node.next[w-'a'] == nil {
				newNode := makeNode()
				node.next[w-'a'] = newNode
			}
			word += string(w)
			node = node.next[w-'a']
		}
		node.word = word
		node.finished = true
	}

	res := ""
	words := strings.Split(sentence, " ")
	for pos, word := range words {
		node := root
		p := 0
		for node != nil && p <= len(word) {
			if node.finished {
				res += node.word
				break
			} else if p < len(word) && node.next[word[p]-'a'] == nil {
				res += word
				break
			} else if p == len(word) {
				res += word
				break
			} else {
				node = node.next[word[p]-'a']
				p++
			}
		}
		if pos != len(words)-1 {
			res += " "
		}
	}

	return res
}

// func main() {
// 	dict := []string{"a", "aa", "aaa", "aaaa"}
// 	sentence := "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"
// 	fmt.Println(replaceWords(dict, sentence))
// }
