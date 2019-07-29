package main

/*
 * @lc app=leetcode id=461 lang=golang
 *
 * [461] Hamming Distance
 */
func hammingDistance(x int, y int) int {
	c := x ^ y
	count := 0
	for c != 0 {
		if c%2 == 1 {
			count++
		}
		c = c >> 1
	}
	return count
}

// func main() {
// 	fmt.Println(hammingDistance(1, 4))
// }
