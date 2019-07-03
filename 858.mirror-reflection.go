package main

/*
 * @lc app=leetcode id=858 lang=golang
 *
 * [858] Mirror Reflection
 */
func mirrorReflection(p int, q int) int {
	if p%q == 0 {
		if p/q%2 == 0 {
			return 2
		}
		return 1
	}

	a := p
	b := q
	tmp := 0
	for b != 0 {
		tmp = a % b
		a = b
		b = tmp
	}
	p = p / a
	q = q / a

	for i := 2; i <= 1000; i++ {
		if q*i%p == 0 {
			if i%2 == 0 {
				return 2
			} else if q%2 == 0 {
				return 0
			} else {
				return 1
			}
		}
	}
	return -1
}

// func main() {
// 	fmt.Println(mirrorReflection(500, 1000))
// }
