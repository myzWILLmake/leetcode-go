package main

/*
 * @lc app=leetcode id=638 lang=golang
 *
 * [638] Shopping Offers
 */
func shoppingOffers(price []int, special [][]int, needs []int) int {
	var checkFunc func([]int) int
	checkFunc = func(currNeeds []int) int {
		min := 0
		for idx, need := range currNeeds {
			min += need * price[idx]
		}

		for _, s := range special {
			for idx, val := range s {
				if idx != len(s)-1 {
					if val > currNeeds[idx] {
						break
					}
				} else {
					for i := 0; i < len(s)-1; i++ {
						currNeeds[i] -= s[i]
					}

					tmp := checkFunc(currNeeds)
					tmp += s[idx]
					if tmp < min {
						min = tmp
					}

					for i := 0; i < len(s)-1; i++ {
						currNeeds[i] += s[i]
					}
				}
			}
		}

		return min
	}

	min := checkFunc(needs)
	return min
}

// func main() {
// 	price := []int{2, 5}
// 	special := [][]int{{3, 0, 5}, {1, 2, 10}}
// 	need := []int{3, 2}
// 	fmt.Println(shoppingOffers(price, special, need))
// }
