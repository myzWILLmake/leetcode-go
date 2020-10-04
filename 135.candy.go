package main

import "fmt"

/*
 * @lc app=leetcode id=135 lang=golang
 *
 * [135] Candy
 */

// @lc code=start
func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	l := make([]int, n)
	r := make([]int, n)
	l[0] = 1
	r[n-1] = 1

	cur := ratings[0]
	for i := 1; i < n; i++ {
		if ratings[i] > cur {
			l[i] = l[i-1] + 1
		} else {
			l[i] = 1
		}
		cur = ratings[i]
	}
	cur = ratings[n-1]
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > cur {
			r[i] = r[i+1] + 1
		} else {
			r[i] = 1
		}
		cur = ratings[i]
	}

	sum := 0
	for i := 0; i < n; i++ {
		if l[i] > r[i] {
			sum += l[i]
		} else {
			sum += r[i]
		}
	}
	return sum
}

// @lc code=end

func main() {
	r := []int{1, 2, 3, 3, 2, 3, 2, 1}
	fmt.Println(candy(r))
}
