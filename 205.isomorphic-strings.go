package main

/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */
func isIsomorphic(s string, t string) bool {
	x := make(map[rune]rune)
	y := make(map[rune]rune)

	for idx, c := range s {
		w := rune(t[idx])
		if x[c] == 0 {
			x[c] = w
		} else {
			if x[c] != w {
				return false
			}
		}

		if y[w] == 0 {
			y[w] = c
		} else {
			if y[w] != c {
				return false
			}
		}
	}

	return true
}

// func main() {
// 	fmt.Println(isIsomorphic("aa", "ab"))
// }
