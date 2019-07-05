package main

/*
 * @lc app=leetcode id=974 lang=golang
 *
 * [974] Subarray Sums Divisible by K
 */

// O(N*N) TLE
// func subarraysDivByK(A []int, K int) int {
// 	ans := 0
// 	n := len(A)
// 	if A[0]%K == 0 {
// 		ans++
// 	}
// 	for i := 1; i < n; i++ {
// 		A[i] += A[i-1]
// 		if A[i]%K == 0 {
// 			ans++
// 		}
// 		for j := 0; j < i; j++ {
// 			if (A[i]-A[j])%K == 0 {
// 				ans++
// 			}
// 		}
// 	}
// 	return ans
// }

func subarraysDivByK(A []int, K int) int {
	n := len(A)
	count := make([]int, K)
	ans := 0
	sum := 0
	count[sum] = 1
	for i := 1; i <= n; i++ {
		sum = ((A[i-1]+sum)%K + K) % K
		count[sum]++
	}

	for _, v := range count {
		if v > 1 {
			ans += v * (v - 1) / 2
		}
	}

	return ans
}

// func main() {
// 	A := []int{4, 5, 0, -2, -3, 1}
// 	K := 5
// 	fmt.Println(subarraysDivByK(A, K))
// }
