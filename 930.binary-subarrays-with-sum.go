package main

/*
 * @lc app=leetcode id=930 lang=golang
 *
 * [930] Binary Subarrays With Sum
 */
func numSubarraysWithSum(A []int, S int) int {
	n := len(A)
	leftZeroCnt := make([]int, n)
	rightZeroCnt := make([]int, n)
	onePos := make([]int, 0)

	cnt := 0
	for i := 0; i < n; i++ {
		if A[i] == 1 {
			leftZeroCnt[i] = cnt
			cnt = 0
			onePos = append(onePos, i)
		} else {
			cnt++
		}
	}

	cnt = 0
	for i := n - 1; i >= 0; i-- {
		if A[i] == 1 {
			rightZeroCnt[i] = cnt
			cnt = 0
		} else {
			cnt++
		}
	}

	if S == 0 {
		if len(onePos) == 0 {
			return n * (n + 1) / 2
		}

		lzc := leftZeroCnt[onePos[0]]
		rzc := rightZeroCnt[onePos[0]]
		ans := lzc*(lzc+1)/2 + rzc*(rzc+1)/2
		for i := 1; i < len(onePos); i++ {
			rzc = rightZeroCnt[onePos[i]]
			ans += rzc * (rzc + 1) / 2
		}
		return ans
	}

	if S == 1 {
		ans := 0
		for i := 0; i < len(onePos); i++ {
			ans += leftZeroCnt[onePos[i]] + rightZeroCnt[onePos[i]]
		}
		return ans + len(onePos)
	}

	ans := 0
	cnt = 0
	for i := S - 1; i < len(onePos); i++ {
		left := i - S + 1
		ans += (leftZeroCnt[onePos[left]] + 1) * (rightZeroCnt[onePos[i]] + 1)
	}
	return ans
}

// func main() {
// 	A := []int{0, 0, 0, 0, 0}
// 	S := 0
// 	fmt.Println(numSubarraysWithSum(A, S))
// }
