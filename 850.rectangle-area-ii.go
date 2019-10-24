package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=850 lang=golang
 *
 * [850] Rectangle Area II
 */
// @lc code=start

type Rect struct {
	data []int
	no   int
}

type Event struct {
	no  int
	val bool
	y   int
}

func getX(active map[int]bool, rects []*Rect) int {
	total := 0
	x1 := -1
	x2 := -1
	for i := 0; i < len(rects); i++ {
		if !active[i] {
			continue
		}
		rect := rects[i]
		if x1 == -1 {
			x1 = rect.data[0]
			x2 = rect.data[2]
		} else {
			if rect.data[0] < x2 && rect.data[2] > x2 {
				x2 = rect.data[2]
			} else if rect.data[0] >= x2 {
				total += x2 - x1
				x1 = rect.data[0]
				x2 = rect.data[2]
			}
		}
	}

	if x1 != -1 {
		total += x2 - x1
	}

	return total
}

func rectangleArea(rectangles [][]int) int {
	rects := make([]*Rect, len(rectangles))
	for i := 0; i < len(rectangles); i++ {
		r := &Rect{rectangles[i], 0}
		rects[i] = r
	}

	sort.SliceStable(rects, func(i, j int) bool {
		return rects[i].data[0] < rects[j].data[0]
	})

	events := []*Event{}

	for i := 0; i < len(rects); i++ {
		rects[i].no = i
		et := &Event{i, true, rects[i].data[1]}
		ef := &Event{i, false, rects[i].data[3]}
		events = append(events, et)
		events = append(events, ef)
	}

	sort.SliceStable(events, func(i, j int) bool {
		return events[i].y < events[j].y
	})

	active := map[int]bool{}

	ans := 0
	y := 0
	for _, e := range events {
		x := getX(active, rects)
		ans += x * (e.y - y)
		y = e.y
		if e.val {
			active[e.no] = true
		} else {
			active[e.no] = false
		}

		ans %= 1000000007
	}

	return ans
}

// func main() {
// 	rect := [][]int{{46, 84, 81, 97}, {17, 14, 21, 52}, {56, 7, 89, 34}, {21, 6, 75, 93}, {5, 13, 30, 19}}
// 	fmt.Println(rectangleArea(rect))
// }

// @lc code=end
