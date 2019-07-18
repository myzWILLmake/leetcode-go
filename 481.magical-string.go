package main

/*
 * @lc app=leetcode id=481 lang=golang
 *
 * [481] Magical String
 */
func magicalString(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	m := make([]int, n)
	m[0] = 1
	m[1] = 2
	lastIdx := 0
	currIdx := 0
	currLen := 1
	currNum := 1
	ans := 1
	for currIdx < n-1 {
		if currLen == m[lastIdx] {
			currNum = 3 - currNum
			lastIdx++
			currLen = 1
			currIdx++
			m[currIdx] = currNum
		} else {
			currIdx++
			currLen++
			m[currIdx] = currNum
		}
		if currNum == 1 {
			ans++
		}
	}
	return ans
}

// func main() {
// 	fmt.Println(magicalString(6))
// }
