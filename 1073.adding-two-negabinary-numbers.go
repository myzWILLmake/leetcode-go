package main

/*
 * @lc app=leetcode id=1073 lang=golang
 *
 * [1073] Adding Two Negabinary Numbers
 */

func addNegabinary(arr1 []int, arr2 []int) []int {

	reverse := func(s []int) []int {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		return s
	}

	arr1 = reverse(arr1)
	arr2 = reverse(arr2)
	n := len(arr1)
	m := len(arr2)
	l := m
	pos := 0
	if n > m {
		l = n
	}

	ans := make([]int, l+2)
	for i := 0; i < l+2; i++ {
		x := 0
		y := 0
		if i < n {
			x = arr1[i]
		}
		if i < m {
			y = arr2[i]
		}
		ans[i] = ans[i] + x + y
		if ans[i] < 0 {
			cnt := -(ans[i])
			ans[i] = 0
			for j := 0; j < cnt; j++ {
				ans[i]++
				ans[i+1]++
			}
		} else {
			for ans[i] > 1 {
				ans[i] = ans[i] - 2
				ans[i+1]--
			}
		}
		if ans[i] > 0 {
			pos = i
		}
	}

	return reverse(ans[:pos+1])
}

// func main() {
// 	arr1 := []int{1, 1, 0, 1}
// 	arr2 := []int{1}
// 	fmt.Println(addNegabinary(arr1, arr2))
// }
