package main

/*
 * @lc app=leetcode id=788 lang=golang
 *
 * [788] Rotated Digits
 */
func rotatedDigits(N int) int {
	ans := 0
	for i := 1; i <= N; i++ {
		isValid := true
		onlyRotateToSelf := true
		num := i
		for num > 0 {
			digit := num % 10
			num = num / 10
			switch digit {
			case 3, 4, 7:
				isValid = false
				break
			case 2, 5, 6, 9:
				onlyRotateToSelf = false
			}
		}
		if isValid && !onlyRotateToSelf {
			ans++
		}
	}
	return ans
}

// func main() {
// 	fmt.Println(rotatedDigits(10000))
// }
