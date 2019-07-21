package main

/*
 * @lc app=leetcode id=931 lang=golang
 *
 * [931] Minimum Falling Path Sum
 */
func minFallingPathSum(A [][]int) int {
	n := len(A)
	m := len(A[0])
	f := make([][]int, n)
	f[0] = make([]int, m)
	for i := 0; i < m; i++ {
		f[0][i] = A[0][i]
	}

	for i := 1; i < n; i++ {
		f[i] = make([]int, m)
		for j := 0; j < m; j++ {
			min := -1
			if j > 0 {
				if min == -1 || f[i-1][j-1] < min {
					min = f[i-1][j-1]
				}
			}

			if min == -1 || f[i-1][j] < min {
				min = f[i-1][j]
			}

			if j < m-1 {
				if min == -1 || f[i-1][j+1] < min {
					min = f[i-1][j+1]
				}
			}

			f[i][j] = min + A[i][j]
		}
	}

	ans := f[n-1][0]
	for j := 1; j < m; j++ {
		if f[n-1][j] < ans {
			ans = f[n-1][j]
		}
	}
	return ans
}

// func main() {
// 	A := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
// 	fmt.Println(minFallingPathSum(A))
// }
