package main

/*
 * @lc app=leetcode id=89 lang=golang
 *
 * [89] Gray Code
 */
func grayCode(n int) []int {
	length := 1 << uint(n)
	codes := make([]int, length)
	codes[0] = 0
	for i := 1; i <= n; i++ {
		interval := 1 << uint(i-1)
		start := 1<<uint(i-1) - 1
		for j, pos := start, start+1; j >= 0; j, pos = j-1, pos+1 {
			codes[pos] = codes[j] + interval
		}
	}

	return codes
}

// func main() {
// 	fmt.Println(grayCode(4))
// }
