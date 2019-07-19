package main

/*
 * @lc app=leetcode id=927 lang=golang
 *
 * [927] Three Equal Parts
 */
func threeEqualParts(A []int) []int {
	n := len(A)
	m := make([]int, n)
	count := 0
	for i := 0; i < n; i++ {
		if A[i] == 1 {
			m[count] = i
			count++
		}
	}
	if count%3 != 0 {
		return []int{-1, -1}
	}

	if count == 0 {
		return []int{0, 2}
	}

	len := count / 3
	for i := 1; i < len; i++ {
		diff1 := m[i] - m[i-1]
		diff2 := m[i+len] - m[i+len-1]
		diff3 := m[i+2*len] - m[i+2*len-1]
		if diff1 != diff2 || diff2 != diff3 || diff1 != diff3 {
			return []int{-1, -1}
		}
	}

	zeroLen := n - 1 - m[count-1]
	if m[len]-m[len-1]-1 < zeroLen || m[len*2]-m[len*2-1]-1 < zeroLen {
		return []int{-1, -1}
	}

	i := m[len-1] + zeroLen
	j := m[len*2-1] + zeroLen + 1

	return []int{i, j}
}

// func main() {
// 	A := []int{1, 1, 0, 0, 1}
// 	fmt.Println(threeEqualParts(A))
// }
