package main

/*
 * @lc app=leetcode id=940 lang=golang
 *
 * [940] Distinct Subsequences II
 */
func distinctSubseqII(S string) int {
	idxMap := make(map[byte]int)
	ans := make([]int, len(S))

	ans[0] = 1
	for idx := 1; idx < len(S); idx++ {
		c := S[idx]
		if idxMap[c] == 0 && c != S[0] {
			ans[idx] = ans[idx-1]*2%1000000007 + 1
		} else {
			if idxMap[c] == 0 {
				ans[idx] = ans[idx-1] * 2 % 1000000007
			} else {
				ans[idx] = (ans[idx-1]*2%1000000007 - ans[idxMap[c]-1]%1000000007 + 1000000007) % 1000000007
			}
		}
		idxMap[c] = idx
	}

	return ans[len(S)-1]
}

// func main() {
// 	fmt.Println(distinctSubseqII("aaa"))
// }
