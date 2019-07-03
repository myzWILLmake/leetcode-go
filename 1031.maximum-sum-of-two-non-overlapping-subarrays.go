package main

/*
 * @lc app=leetcode id=1031 lang=golang
 *
 * [1031] Maximum Sum of Two Non-Overlapping Subarrays
 */
// Time Complexity: O(N*N)
// func maxSumTwoNoOverlap(A []int, L int, M int) int {
// 	n := len(A)
// 	sumL := make([]int, n)
// 	for i := 0; i < L; i++ {
// 		sumL[0] += A[i]
// 	}

// 	ans := 0
// 	for i := 0; i < n-L+1; i++ {
// 		if i > 0 {
// 			sumL[i] = sumL[i-1] - A[i-1] + A[i+L-1]
// 		}
// 		ansM := -1
// 		for j := 0; j < n-M+1; j++ {
// 			if (j >= i && j <= i+L-1) || (j+M-1 >= i && j+M-1 <= i+L-1) ||
// 				(i >= j && i <= j+M-1) || (i+L-1 >= j && i+L-1 <= j+M-1) {
// 				continue
// 			}
// 			sumM := 0
// 			for k := j; k <= j+M-1; k++ {
// 				sumM += A[k]
// 			}
// 			if ansM == -1 || sumM > ansM {
// 				ansM = sumM
// 			}
// 		}
// 		if ansM == -1 {
// 			continue
// 		} else if ansM+sumL[i] > ans {
// 			ans = ansM + sumL[i]
// 		}
// 	}

// 	return ans
// }

// dp Time Complexity O(N)
// ans = max(ans, A[i]-A[i-L]+b[i-L], A[i]-A[i-M]+a[i-M])
func maxSumTwoNoOverlap(A []int, L int, M int) int {
	n := len(A)
	a := make([]int, n)
	b := make([]int, n)
	ans := 0
	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}

	for i := range A {
		if i > 0 {
			A[i] += A[i-1]
		}
	}

	for i := 0; i < n; i++ {
		if i < L {
			a[i] = A[i]
		} else {
			a[i] = max(a[i-1], A[i]-A[i-L])
		}

		if i < M {
			b[i] = A[i]
		} else {
			b[i] = max(b[i-1], A[i]-A[i-M])
		}

		if i < L+M {
			ans = A[i]
		} else {
			tmp := max(A[i]-A[i-L]+b[i-L], A[i]-A[i-M]+a[i-M])
			ans = max(ans, tmp)
		}
	}
	return ans

}

// func main() {
// 	a := []int{4, 5, 14, 16, 16, 20, 7, 13, 8, 15}
// 	fmt.Println(maxSumTwoNoOverlap(a, 3, 5))
// }
