package main

/*
 * @lc app=leetcode id=446 lang=golang
 *
 * [446] Arithmetic Slices II - Subsequence
 */
func numberOfArithmeticSlices(A []int) int {
	//f[i][j] = sum(f[j][k]) i-j=j-k
	n := len(A)
	count := make([][]int, n)
	for i := 0; i < n; i++ {
		count[i] = make([]int, n)
	}

	ans := 0
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			sub := A[i] - A[j]
			tmp := 0
			for k := j + 1; k < n; k++ {
				if A[j]-A[k] == sub {
					tmp += count[j][k] + 1
				}
			}
			count[i][j] = tmp
			ans += tmp
		}
	}

	return ans
}

// func main() {
// 	nums := []int{0, 1, 2, 2, 2}
// 	fmt.Println(numberOfArithmeticSlices(nums))
// }
