package main

import "fmt"

/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start
type Trie struct {
	end  bool
	next []*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	it := Trie{}
	it.end = false
	it.next = make([]*Trie, 26)
	return it
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	it := this
	for _, c := range word {
		if it.next[c-'a'] != nil {
			it = it.next[c-'a']
		} else {
			newit := Constructor()
			it.next[c-'a'] = &newit
			it = &newit
		}
	}

	it.end = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	it := this
	for _, c := range word {
		if it.next[c-'a'] != nil {
			it = it.next[c-'a']
		} else {
			return false
		}
	}

	if !it.end {
		return false
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	it := this
	for _, c := range prefix {
		if it.next[c-'a'] != nil {
			it = it.next[c-'a']
		} else {
			return false
		}
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

func main() {
	obj := Constructor()
	obj.Insert("apple")
	fmt.Println(obj.Search("apple"))
	fmt.Println(obj.Search("app"))
	fmt.Println(obj.StartsWith("app"))
	obj.Insert("app")
	fmt.Println(obj.Search("app"))
}
