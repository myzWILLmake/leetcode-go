/*
 * @lc app=leetcode id=994 lang=golang
 *
 * [994] Rotting Oranges
 */
package main

func orangesRotting(grid [][]int) int {
	count := 0
	newRotten := true
	for newRotten {
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == 3 {
					grid[i][j] = 2
				}
			}
		}

		newRotten = false
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == 2 {
					if i > 0 && grid[i-1][j] == 1 {
						newRotten = true
						grid[i-1][j] = 3
					}
					if j > 0 && grid[i][j-1] == 1 {
						newRotten = true
						grid[i][j-1] = 3
					}
					if i < len(grid)-1 && grid[i+1][j] == 1 {
						newRotten = true
						grid[i+1][j] = 3
					}
					if j < len(grid[i])-1 && grid[i][j+1] == 1 {
						newRotten = true
						grid[i][j+1] = 3
					}
				}
			}
		}
		if newRotten {
			count++
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return count
}

// func main() {
// 	grid := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
// 	fmt.Println(orangesRotting(grid))
// }
