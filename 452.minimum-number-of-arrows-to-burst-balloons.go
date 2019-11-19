package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=452 lang=golang
 *
 * [452] Minimum Number of Arrows to Burst Balloons
 */

// @lc code=start
type Point struct {
	start int
	end   int
}

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	p := make([]*Point, len(points))

	for i := 0; i < len(points); i++ {
		point := &Point{}
		point.start = points[i][0]
		point.end = points[i][1]
		p[i] = point
	}

	sort.SliceStable(p, func(i, j int) bool {
		return p[i].start < p[j].start
	})

	ans := 1
	end := p[0].end

	for i := 0; i < len(p); i++ {
		point := p[i]
		if point.start > end {
			ans++
			end = point.end
		} else {
			if point.end < end {
				end = point.end
			}
		}
	}

	return ans
}

// @lc code=end
func main() {
	points := [][]int{{1, 3}, {2, 4}, {3, 5}, {4, 6}, {5, 7}}
	fmt.Println(findMinArrowShots(points))
}
