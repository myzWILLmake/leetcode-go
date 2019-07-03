package main

import (
	"math"
)

// Find the median of the two sorted arrays.
// The overall run time complexity should be O(log (m+n)).
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var n1, n2 []int
	var m, n, iMin, iMax, halfLen int

	m = len(nums1)
	n = len(nums2)

	if m > n {
		n2 = nums1[:]
		n1 = nums2[:]
		m = n
		n = len(n2)
	} else {
		n1 = nums1[:]
		n2 = nums2[:]
	}

	iMin = 0
	iMax = m
	halfLen = (m + n + 1) / 2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < iMax && n2[j-1] > n1[i] {
			iMin = i + 1
		} else if i > iMin && n1[i-1] > n2[j] {
			iMax = i - 1
		} else {
			var maxLeft, minRight float64
			maxLeft = 0
			if i == 0 {
				maxLeft = float64(n2[j-1])
			} else if j == 0 {
				maxLeft = float64(n1[i-1])
			} else {
				maxLeft = math.Max(float64(n1[i-1]), float64(n2[j-1]))
			}

			if (m+n)%2 == 1 {
				return maxLeft
			}

			minRight = 0
			if i == m {
				minRight = float64(n2[j])
			} else if j == n {
				minRight = float64(n1[i])
			} else {
				minRight = math.Min(float64(n1[i]), float64(n2[j]))
			}

			return (maxLeft + minRight) / 2.0
		}
	}

	return 0
}

// func main() {
// 	nums1 := []int{1, 3}
// 	nums2 := []int{2}
// 	res := findMedianSortedArrays(nums1, nums2)
// 	fmt.Println(res)
// }
