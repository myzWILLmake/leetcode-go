package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	headPos := 0
	res := 1
	letterMap := make(map[rune]int)
	flagMap := make(map[rune]bool)

	for pos, ch := range s {
		if flagMap[ch] {
			length := pos - headPos
			if length > res {
				res = length
			}
			if headPos < letterMap[ch]+1 {
				headPos = letterMap[ch] + 1
			}
		} else {
			flagMap[ch] = true
		}
		letterMap[ch] = pos
	}

	if len(s)-headPos > res {
		res = len(s) - headPos
	}

	return res
}

// func main() {
// 	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
// }
