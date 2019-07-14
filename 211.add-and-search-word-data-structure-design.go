package main

/*
 * @lc app=leetcode id=211 lang=golang
 *
 * [211] Add and Search Word - Data structure design
 */
type TrieNode struct {
	finished bool
	next     []*TrieNode
}

type WordDictionary struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	root := TrieNode{}
	root.next = make([]*TrieNode, 26)
	w := WordDictionary{&root}
	return w
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, c := range word {
		if node.next[c-'a'] != nil {
			node = node.next[c-'a']
		} else {
			newNode := TrieNode{}
			newNode.next = make([]*TrieNode, 26)
			node.next[c-'a'] = &newNode
			node = &newNode
		}
	}
	node.finished = true
}

func nodeSearch(idx int, word string, node *TrieNode) bool {
	for i := idx; i < len(word); i++ {
		c := word[i]
		if c == '.' {
			check := false
			for j := 0; j < 26; j++ {
				if node.next[j] != nil {
					check = check || nodeSearch(i+1, word, node.next[j])
					if check {
						return true
					}
				}
			}
			return false
		} else {
			if node.next[c-'a'] != nil {
				node = node.next[c-'a']
			} else {
				return false
			}
		}
	}

	if node.finished {
		return true
	}
	return false
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return nodeSearch(0, word, this.root)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

// func main() {
// 	obj := Constructor()
// 	obj.AddWord("bad")
// 	obj.AddWord("dad")
// 	obj.AddWord("mad")
// 	fmt.Println(obj.Search("pad"))
// 	fmt.Println(obj.Search("bad"))
// 	fmt.Println(obj.Search(".ad"))
// 	fmt.Println(obj.Search("b.."))
// 	fmt.Println(obj.Search("..."))
// }
