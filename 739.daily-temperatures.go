package main

/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */
func dailyTemperatures(T []int) []int {
	type Tmp struct {
		Val int
		day int
	}
	n := len(T)
	stack := make([]Tmp, n)
	res := make([]int, n)
	top := -1
	for i := 0; i < n; i++ {
		newTmp := Tmp{T[i], i}
		if top >= 0 {
			topTmp := stack[top]
			for topTmp.Val < newTmp.Val && top >= 0 {
				top--
				res[topTmp.day] = newTmp.day - topTmp.day
				if top >= 0 {
					topTmp = stack[top]
				}
			}
		}
		top++
		stack[top] = newTmp
	}

	return res
}

// func main() {
// 	tmp := []int{73, 74, 75, 71, 69, 72, 76, 73}
// 	fmt.Println(dailyTemperatures(tmp))
// }
