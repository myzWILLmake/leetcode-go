package main

/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */
func getRow(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	tmp := make([]int, rowIndex+1)

	for i := 0; i <= rowIndex; i++ {
		tmp[0] = 1
		tmp[rowIndex] = 1
		if rowIndex > 1 {
			for j := 1; j <= rowIndex-1; j++ {
				tmp[j] = ans[j-1] + ans[j]
			}
		}
		swap := ans
		ans = tmp
		tmp = swap
	}
	return ans
}

// func main() {
// 	fmt.Println(getRow(0))
// }
