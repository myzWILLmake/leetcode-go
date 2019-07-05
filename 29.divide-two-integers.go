package main

/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	binaryDiv := func(a, b int) int64 {
		var ans int64
		ans = 0
		for i := 31; i >= 0; i-- {
			if a>>uint(i) >= b {
				ans += 1 << uint(i)
				a -= b << uint(i)
			}
		}
		return ans
	}

	var ans int64
	ans = 0
	if dividend < 0 {
		if divisor < 0 {
			ans = binaryDiv(-dividend, -divisor)
		} else {
			ans = -binaryDiv(-dividend, divisor)
		}
	} else if divisor < 0 {
		ans = -binaryDiv(dividend, -divisor)
	} else {
		ans = binaryDiv(dividend, divisor)
	}
	if ans > (1<<31)-1 || ans < -(1<<31) {
		return (1 << 31) - 1
	}
	return int(ans)
}

// func main() {
// 	fmt.Println(divide(-2147483648, -1))
// 	fmt.Println(-2147483648 / -1)
// }
