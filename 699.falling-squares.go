package main

/*
 * @lc app=leetcode id=699 lang=golang
 *
 * [699] Falling Squares
 */

type Data struct {
	idx    int
	pos    int
	length int
	height int
}

func fallingSquares(positions [][]int) []int {
	n := len(positions)
	allData := make([]*Data, n)

	ans := 0
	allAns := make([]int, n)
	for idx, position := range positions {
		data := Data{}
		data.idx = idx
		data.pos = position[0]
		data.length = position[1]

		maxHeight := 0
		for i := 0; i < idx; i++ {
			tmpData := allData[i]
			if tmpData.pos >= data.pos && tmpData.pos < data.pos+data.length ||
				tmpData.pos+tmpData.length > data.pos && tmpData.pos+tmpData.length <= data.pos+data.length ||
				tmpData.pos <= data.pos && tmpData.pos+tmpData.length >= data.pos+data.length {
				if tmpData.height > maxHeight {
					maxHeight = tmpData.height
				}
			}
		}

		data.height = maxHeight + data.length
		allData[idx] = &data
		if data.height > ans {
			ans = data.height
		}
		allAns[idx] = ans
	}

	return allAns
}

// func main() {
// 	input := [][]int{{1, 5}, {2, 2}, {7, 5}}
// 	fmt.Println(fallingSquares(input))
// }
