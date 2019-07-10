package main

/*
 * @lc app=leetcode id=313 lang=golang
 *
 * [313] Super Ugly Number
 */
func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}

	ugly := make([]int, n)
	pointers := make([]int, len(primes))

	ugly[0] = 1

	for i := 1; i < n; i++ {

		minNum := -1
		choose := []int{}
		for idx, num := range primes {
			tmp := ugly[pointers[idx]] * num
			if tmp > 1<<32 {
				continue
			}
			if minNum == -1 || tmp < minNum {
				minNum = tmp
				choose = []int{idx}
			} else if tmp == minNum {
				choose = append(choose, idx)
			}
		}

		ugly[i] = minNum
		for _, idx := range choose {
			pointers[idx]++
		}
	}

	return ugly[n-1]
}

// func main() {
// 	n := 12
// 	primes := []int{2, 7, 13, 19}
// 	fmt.Println(nthSuperUglyNumber(n, primes))
// }
