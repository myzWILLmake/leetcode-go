package main

/*
 * @lc app=leetcode id=810 lang=golang
 *
 * [810] Chalkboard XOR Game
 */
func xorGame(nums []int) bool {
	curr := 0
	ans := true
	for i := 0; i < len(nums); i++ {
		curr = 0
		for _, num := range nums {
			if num != -1 {
				curr ^= num
			}
		}
		if curr == 0 {
			return ans
		}

		selected := false
		for idx, num := range nums {
			if num != -1 && num^curr != 0 {
				nums[idx] = -1
				selected = true
				break
			}
		}

		if !selected {
			return !ans
		}
		ans = !ans

	}
	return ans
}

// func main() {
// 	nums := []int{1, 1, 2}
// 	fmt.Println(xorGame(nums))
// }
