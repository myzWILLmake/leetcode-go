package main

/*
 * @lc app=leetcode id=805 lang=golang
 *
 * [805] Split Array With Same Average
 */

func splitArraySameAverage(A []int) bool {
	n := len(A)
	half := n / 2
	total := 0
	for i := 0; i < n; i++ {
		total += A[i]
	}

	isPossible := false
	for i := 1; i <= half; i++ {
		if total*i%n == 0 {
			isPossible = true
			break
		}
	}
	if !isPossible {
		return false
	}

	sumset := make([]map[int]bool, half+1)
	for i := 0; i <= half; i++ {
		sumset[i] = make(map[int]bool)
	}
	sumset[0][0] = true
	for _, num := range A {
		for i := half; i > 0; i-- {
			for k := range sumset[i-1] {
				sumset[i][k+num] = true
			}
		}
	}

	for i := 1; i <= half; i++ {
		if total*i%n == 0 && sumset[i][total*i/n] {
			return true
		}
	}

	return false
}

// func main() {
// 	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	fmt.Println(splitArraySameAverage(s))
// }
