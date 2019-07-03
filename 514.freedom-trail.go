package main

/*
 * @lc app=leetcode id=514 lang=golang
 *
 * [514] Freedom Trail
 */

func findRotateSteps(ring string, key string) int {
	var abs = func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	ringLetters := make([][]int, 26)
	for pos, w := range ring {
		ringLetters[w-'a'] = append(ringLetters[w-'a'], pos)
	}

	n := len(ring)
	m := len(key)
	res := make([][]int, m)
	startKeyLetterPos := ringLetters[key[0]-'a']
	res[0] = make([]int, len(startKeyLetterPos))
	for i := 0; i < len(res[0]); i++ {
		res[0][i] = min(startKeyLetterPos[i], n-startKeyLetterPos[i]) + 1
	}

	for i := 1; i < m; i++ {
		keyLetterPos := ringLetters[key[i]-'a']
		res[i] = make([]int, len(keyLetterPos))
		for j := 0; j < len(keyLetterPos); j++ {
			minStep := -1
			lastLetterPos := ringLetters[key[i-1]-'a']
			for k := 0; k < len(lastLetterPos); k++ {
				tmp := res[i-1][k] + min(abs(keyLetterPos[j]-lastLetterPos[k]), n-abs(keyLetterPos[j]-lastLetterPos[k]))
				if minStep == -1 || tmp < minStep {
					minStep = tmp
				}
			}
			res[i][j] = minStep + 1
		}
	}

	finalLettersPos := ringLetters[key[m-1]-'a']
	minRes := -1
	for j := 0; j < len(finalLettersPos); j++ {
		if minRes == -1 || res[m-1][j] < minRes {
			minRes = res[m-1][j]
		}
	}
	return minRes
}

// func main() {
// 	ring := "pqwcx"
// 	key := "cpqwx"
// 	fmt.Println(findRotateSteps(ring, key))
// }
