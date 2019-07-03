package main

/*
 * @lc app=leetcode id=840 lang=golang
 *
 * [840] Magic Squares In Grid
 */
func numMagicSquaresInside(grid [][]int) int {
	ans := 0

	n := len(grid)
	m := len(grid[0])

	if n < 3 || m < 3 {
		return 0
	}

	for i := 0; i <= n-3; i++ {
		for j := 0; j <= m-3; j++ {
			check := true
			checkNum := make([]bool, 10)

			sum := grid[i][j] + grid[i][j+1] + grid[i][j+2]
			for x := i; x < i+3; x++ {
				tmpSum := 0
				for y := j; y < j+3; y++ {
					if grid[x][y] > 9 || grid[x][y] < 1 {
						check = false
						break
					}
					tmpSum += grid[x][y]
					checkNum[grid[x][y]] = true
				}
				if tmpSum != sum {
					check = false
					break
				}
			}

			for i := 1; i < 10; i++ {
				if !checkNum[i] {
					check = false
					break
				}
			}

			if !check {
				continue
			}

			for y := j; y < j+3; y++ {
				tmpSum := 0
				for x := i; x < i+3; x++ {
					tmpSum += grid[x][y]
				}
				if tmpSum != sum {
					check = false
					break
				}
			}

			sumDia := grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2]
			if sumDia != sum {
				check = false
			}

			sumDia = grid[i+2][j] + grid[i+1][j+1] + grid[i][j+2]
			if sumDia != sum {
				check = false
			}

			if check {
				ans++
			}
		}
	}

	return ans
}

// func main() {
// 	grid := [][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}
// 	fmt.Println(numMagicSquaresInside(grid))
// }
