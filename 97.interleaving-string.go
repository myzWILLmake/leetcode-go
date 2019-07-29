package main

/*
 * @lc app=leetcode id=97 lang=golang
 *
 * [97] Interleaving String
 */

func checkFunc(s1, s2, s3 string, i, j, k int) bool {
	if i == len(s1) && j == len(s2) && k == len(s3) {
		return true
	}

	if i < len(s1) && s1[i] == s3[k] {
		if checkFunc(s1, s2, s3, i+1, j, k+1) {
			return true
		}
	}

	if j < len(s2) && s2[j] == s3[k] {
		if checkFunc(s1, s2, s3, i, j+1, k+1) {
			return true
		}
	}

	return false
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}

	return checkFunc(s1, s2, s3, 0, 0, 0)
}

// func main() {
// 	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
// }
