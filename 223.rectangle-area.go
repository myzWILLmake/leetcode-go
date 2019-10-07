/*
 * @lc app=leetcode id=223 lang=golang
 *
 * [223] Rectangle Area
 */
package main

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	total := (C-A)*(D-B) + (G-E)*(H-F)
	x := 0
	y := 0
	if A <= E && E <= C {
		if G < C {
			x = G - E
		} else {
			x = C - E
		}
	}

	if E <= A && A <= G {
		if C < G {
			x = C - A
		} else {
			x = G - A
		}
	}

	if B <= F && F <= D {
		if H < D {
			y = H - F
		} else {
			y = D - F
		}
	}

	if F <= B && B <= H {
		if D < H {
			y = D - B
		} else {
			y = H - B
		}
	}

	total = total - x*y
	return total
}

// func main() {
// 	fmt.Println(computeArea(-2, -2, 2, 2, -2, -2, 2, 2))
// }
