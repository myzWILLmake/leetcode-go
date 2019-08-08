package main

/*
 * @lc app=leetcode id=419 lang=golang
 *
 * [419] Battleships in a Board
 */
func countBattleships(board [][]byte) int {
	if len(board) == 0 {
		return 0
	}
	cnt := 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'X' {
				if i == 0 && j == 0 && board[0][0] == 'X' {
					cnt++
				} else {
					if i == 0 {
						if board[i][j-1] != 'X' {
							cnt++
						}
					} else if j == 0 {
						if board[i-1][j] != 'X' {
							cnt++
						}
					} else {
						if board[i-1][j] != 'X' && board[i][j-1] != 'X' {
							cnt++
						}
					}
				}
			}
		}
	}

	return cnt
}

// func main() {
// 	board := [][]byte{{'X', '.', '.', 'X'}, {'.', 'X', 'X', '.'}, {'.', '.', '.', 'X'}}
// 	fmt.Println(countBattleships(board))
// }
