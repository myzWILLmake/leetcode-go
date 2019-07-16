package main

/*
 * @lc app=leetcode id=794 lang=golang
 *
 * [794] Valid Tic-Tac-Toe State
 */
func validTicTacToe(board []string) bool {
	b := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	countX := 0
	countO := 0
	for x, s := range board {
		for y, c := range s {
			if c == 'X' {
				b[x][y] = 1
				countX++
			} else if c == 'O' {
				b[x][y] = -1
				countO++
			}
		}
	}

	if countX-countO < 0 || countX-countO > 1 {
		return false
	}

	winXCount := 0
	winOCount := 0
	for i := 0; i < 3; i++ {
		sum := 0
		for j := 0; j < 3; j++ {
			sum += b[i][j]
		}
		if sum == 3 {
			winXCount++
		} else if sum == -3 {
			winOCount++
		}
	}

	for i := 0; i < 3; i++ {
		sum := 0
		for j := 0; j < 3; j++ {
			sum += b[j][i]
		}
		if sum == 3 {
			winXCount++
		} else if sum == -3 {
			winOCount++
		}
	}

	sum := b[0][0] + b[1][1] + b[2][2]
	if sum == 3 {
		winXCount++
	} else if sum == -3 {
		winOCount++
	}
	sum = b[2][0] + b[1][1] + b[0][2]
	if sum == 3 {
		winXCount++
	} else if sum == -3 {
		winOCount++
	}

	if winXCount > 0 {
		if countO >= countX {
			return false
		}
	}

	if winOCount > 0 {
		if countO != countX {
			return false
		}
	}

	if winXCount > 0 && winOCount > 0 {
		return false
	}

	return true
}

// func main() {
// 	board := []string{"XOX", "O O", "XOX"}
// 	fmt.Println(validTicTacToe(board))
// }
// XXX
// XOO
// OO
