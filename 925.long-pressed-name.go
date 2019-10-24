package main

/*
 * @lc app=leetcode id=925 lang=golang
 *
 * [925] Long Pressed Name
 */

// @lc code=start
func isLongPressedName(name string, typed string) bool {
	posx, posy := 0, 0
	for posx < len(name) && posy < len(typed) {
		c := name[posx]
		if c != typed[posy] {
			return false
		}
		cntx := 1
		for posx+1 < len(name) && name[posx] == name[posx+1] {
			cntx++
			posx++
		}

		cnty := 1
		for posy+1 < len(typed) && typed[posy] == typed[posy+1] {
			cnty++
			posy++
		}

		if cnty < cntx {
			return false
		}

		posx++
		posy++
	}

	if posx == len(name) && posy == len(typed) {
		return true
	}
	return false
}

// func main() {
// 	fmt.Println(isLongPressedName("alex", "aalee"))
// }

// @lc code=end
