package main

/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

// func main() {
// 	s := []byte{'h', 'e', 'l', 'l', 'o'}
// 	reverseString(s)
// 	fmt.Println(s)
// }
