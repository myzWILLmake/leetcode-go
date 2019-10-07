package main

/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */
func firstUniqChar(s string) int {
	pos := make([]int, 26)

	for idx, c := range s {
		if pos[c-'a'] == 0 {
			pos[c-'a'] = idx + 1
		} else {
			pos[c-'a'] = -1
		}
	}

	min := len(s) + 1
	for i := 0; i < 26; i++ {
		if pos[i] != 0 && pos[i] != -1 && pos[i] < min {
			min = pos[i]
		}
	}
	if min == len(s)+1 {
		return -1
	}
	return min - 1
}

// func main() {
// 	fmt.Println(firstUniqChar(""))
// }
