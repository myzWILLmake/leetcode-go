/*
 * @lc app=leetcode id=482 lang=golang
 *
 * [482] License Key Formatting
 */
package main

import (
	"strings"
)

func licenseKeyFormatting(S string, K int) string {
	S = strings.ToUpper(S)
	res := []byte{}
	count := 0

	for i := len(S) - 1; i >= 0; i-- {
		if S[i] != '-' {
			if count == K {
				res = append(res, '-')
				count = 0
			}

			res = append(res, S[i])
			count++
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}

// func main() {
// 	fmt.Println(licenseKeyFormatting("-c", 4))
// }
